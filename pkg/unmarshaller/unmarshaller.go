package unmsarshaller

import (
	"strings"

	godyn "github.com/vedadiyan/godyn/pkg"
	"github.com/vedadiyan/proton/pkg/conversions"
	"github.com/vedadiyan/proton/pkg/functions"
	"github.com/vedadiyan/proton/pkg/helpers"
	"github.com/vedadiyan/proton/pkg/models"
	"github.com/vedadiyan/proton/pkg/options"
	"github.com/vedadiyan/proton/pkg/sentinel"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type unmarshaller struct {
	data map[string]any
}

type ExprCtx *godyn.Godyn[models.ProtonArg]

var context ExprCtx

func init() {
	functions := functions.GetFunctions()
	context = &functions
}

func Unmarshall(data map[string]any, pb proto.Message) error {
	u := unmarshaller{
		data: data,
	}
	return u.parse(data, pb)
}

func (u unmarshaller) parse(data map[string]any, pb proto.Message) error {
	fields := pb.ProtoReflect().Descriptor().Fields()
	len := fields.Len()

	for i := 0; i < len; i++ {
		fd := fields.Get(i)
		err := u.link(context, data, pb, fd)
		if err != nil {
			if err.Error() == sentinel.TargetDoesNotExist().Error() {
				continue
			}
			return err
		}
	}

	return nil
}

func (u unmarshaller) link(exprCtx ExprCtx, data map[string]any, pb proto.Message, fd protoreflect.FieldDescriptor) error {
	if fd.IsList() {
		l, err := u.readList(exprCtx, data, pb, fd)
		if err != nil {
			if err.Error() != sentinel.TargetDoesNotExist().Error() {
				return err
			}
		}
		return u.setList(l, pb, fd)
	}
	switch fd.Kind() {
	case protoreflect.MessageKind:
		{
			d, err := u.readMessage(exprCtx, data, pb, fd)
			if err != nil {
				return err
			}
			return u.setMessage(d, pb, fd)
		}
	default:
		{
			v, err := u.readValue(exprCtx, data, pb, fd)
			if err != nil {
				return err
			}
			return u.setValue(v, pb, fd)
		}
	}
}

func (u unmarshaller) readList(exprCtx ExprCtx, data map[string]any, pb proto.Message, fd protoreflect.FieldDescriptor) ([]any, error) {
	opt := options.GetOptions(fd)
	fieldName := _select(stringNotEmpty, opt.GetSelectAttribute(), helpers.FieldName(fd))
	if fieldName == nil {
		panic("not implemented")
	}

	if opt.HasReduceAttribute() {
		res, err := (*exprCtx).Invoke(models.New(u.data, data), opt.GetReduceAttribute())
		if err != nil {
			return nil, err
		}
		switch r := res.(type) {
		case []any:
			{
				return r, nil
			}
		default:
			{
				return []any{r}, nil
			}
		}
	}

	target := extract(u.data, data, *fieldName)
	if target == nil {
		return nil, sentinel.TargetDoesNotExist()
	}
	if list, ok := target.([]any); ok {
		return list, nil
	}
	if list, ok := target.(map[string]any); ok {
		return []any{list}, nil
	}
	return nil, sentinel.ObjectIsNotArray()
}

func (u unmarshaller) readMessage(exprCtx ExprCtx, data map[string]any, pb proto.Message, fd protoreflect.FieldDescriptor) (protoreflect.ProtoMessage, error) {
	opt := options.GetOptions(fd)
	swappedData, err := dataSwap(opt, u.data, data)
	if err != nil {
		return nil, err
	}
	message := pb.ProtoReflect().Get(fd).Message().Interface().ProtoReflect().New().Interface()
	err = u.parse(swappedData, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (u unmarshaller) readValue(exprCtx ExprCtx, data map[string]any, pb proto.Message, fd protoreflect.FieldDescriptor) (any, error) {
	opt := options.GetOptions(fd)
	if !opt.HasAttributes() {
		fieldName := helpers.FieldName(fd)
		return extract(u.data, data, fieldName), nil
	}
	if opt.HasReduceAttribute() {
		res, err := (*exprCtx).Invoke(models.New(u.data, data), opt.GetReduceAttribute())
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return extract(u.data, data, opt.GetSelectAttribute()), nil
}

func (u unmarshaller) setMessage(message protoreflect.ProtoMessage, pb proto.Message, fd protoreflect.FieldDescriptor) error {

	pb.ProtoReflect().Set(fd, protoreflect.ValueOfMessage(message.ProtoReflect()))
	return nil
}

func (u unmarshaller) setList(list []any, pb proto.Message, fd protoreflect.FieldDescriptor) error {
	if len(list) == 0 {
		return nil
	}

	if _, ok := list[0].(map[string]any); ok {
		return u.setListComplex(list, pb, fd)
	}
	return u.setListPrimitive(list, pb, fd)
}

func (u unmarshaller) setListPrimitive(list []any, pb proto.Message, fd protoreflect.FieldDescriptor) error {
	if list == nil {
		return nil
	}
	ls := pb.ProtoReflect().Mutable(fd).List()
	for _, item := range list {
		if item == nil {
			continue
		}
		ls.Append(protoreflect.ValueOf(item))
	}
	return nil
}

func (u unmarshaller) setListComplex(list []any, pb proto.Message, fd protoreflect.FieldDescriptor) error {
	ls := pb.ProtoReflect().Mutable(fd).List()
	for _, item := range list {
		message := ls.NewElement().Message().Interface()
		err := u.parse(item.(map[string]any), message)
		if err != nil {
			return err
		}
		ls.Append(protoreflect.ValueOf(message.ProtoReflect()))
	}
	return nil
}

func (u unmarshaller) setValue(value any, pb proto.Message, fd protoreflect.FieldDescriptor) error {
	if value == nil {
		return nil
	}
	convert, err := conversions.Convert(fd.Kind(), value)
	if err != nil {
		return err
	}
	pb.ProtoReflect().Set(fd, protoreflect.ValueOf(convert))
	return nil
}

func _select[T any](selector func(t T) bool, items ...T) *T {
	for _, item := range items {
		if selector(item) {
			return &item
		}
	}
	return nil
}

// func mapNotNull(data map[string]any) bool {
// 	return data != nil
// }

func stringNotEmpty(str string) bool {
	return str != ""
}

func dataSwap(opt *options.ProtonOptions, rootData map[string]any, scopedData map[string]any) (map[string]any, error) {
	if opt.HasSelectAttribute() {
		e := extract(rootData, scopedData, opt.GetSelectAttribute())
		if e == nil {
			return nil, nil
		}
		data, ok := e.(map[string]any)
		if !ok {
			return nil, sentinel.InvalidDataType()
		}
		return data, nil
	}
	return rootData, nil
}

func extract(rootData map[string]any, scopedData map[string]any, key string) any {
	if !stringNotEmpty(key) {
		return scopedData
	}

	keys := strings.Split(key, ";")
	for _, key := range keys {
		d := scopedData
		if strings.HasPrefix(key, "$root.") {
			d = rootData
			key = strings.TrimPrefix(key, "$root.")
		}
		segments := strings.Split(key, ".")
		len := len(segments)
		value := d[segments[0]]
		for i := 1; i < len; i++ {
			_map, ok := value.(map[string]any)
			if !ok {
				break
			}
			value = _map[segments[i]]
		}
		if value != nil {
			return value
		}
	}
	return nil
}
