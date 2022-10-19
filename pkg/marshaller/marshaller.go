package marshaller

import (
	"strings"

	"github.com/vedadiyan/proton/pkg/helpers"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Marshall(pb proto.Message, data map[string]any) error {
	return parse(data, pb)
}

func parse(data map[string]any, pb proto.Message) error {
	fields := pb.ProtoReflect().Descriptor().Fields()
	len := fields.Len()

	for i := 0; i < len; i++ {
		fd := fields.Get(i)
		err := link(data, pb, fd)
		if err != nil {
			return err
		}
	}
	return nil
}

func link(data map[string]any, pb proto.Message, fd protoreflect.FieldDescriptor) error {
	fieldName := helpers.FieldName(fd)
	if fd.IsList() {
		ls, err := readList(data, pb, fd)
		if err != nil {
			return err
		}
		setValue(data, ls, fieldName)
	}
	switch fd.Kind() {
	case protoreflect.MessageKind:
		{
			v := readMessage(pb, fd)

			d := make(map[string]any)
			err := parse(d, v)
			if err != nil {
				return err
			}
			setValue(data, d, fieldName)
		}
	default:
		{
			v := readValue(pb, fd)
			setValue(data, v, fieldName)
		}
	}
	return nil
}

func readList(data map[string]any, pb proto.Message, fd protoreflect.FieldDescriptor) ([]any, error) {
	switch fd.Kind() {
	case protoreflect.MessageKind:
		{
			return readListComplex(pb, fd)
		}
	default:
		{
			return readListPrimitive(pb, fd)
		}
	}
}
func readMessage(pb proto.Message, fd protoreflect.FieldDescriptor) protoreflect.ProtoMessage {
	v := pb.ProtoReflect().Get(fd).Message().Interface()
	return v
}
func readListComplex(pb proto.Message, fd protoreflect.FieldDescriptor) ([]any, error) {
	list := make([]any, 0)
	ls := pb.ProtoReflect().Get(fd).List()
	len := ls.Len()

	for i := 0; i < len; i++ {
		item := ls.Get(i).Message()
		d := make(map[string]any)
		err := parse(d, item.Interface())
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

func readListPrimitive(pb proto.Message, fd protoreflect.FieldDescriptor) ([]any, error) {
	list := make([]any, 0)
	ls := pb.ProtoReflect().Get(fd).List()
	len := ls.Len()

	for i := 0; i < len; i++ {
		item := ls.Get(i).Interface()
		list = append(list, item)
	}
	return list, nil
}

func readValue(pb proto.Message, fd protoreflect.FieldDescriptor) any {
	return pb.ProtoReflect().Get(fd).Interface()
}

func setValue(data map[string]any, value any, fieldName string) {
	segments := strings.Split(strings.Split(fieldName, ";")[0], ".")
	len := len(segments)
	v := data
	for i := 0; i < len-1; i++ {
		if v[segments[i]] != nil {
			v = v[segments[i]].(map[string]any)
			continue
		}
		v[segments[i]] = make(map[string]any)
		v = v[segments[i]].(map[string]any)
	}
	v[segments[len-1]] = value
}
