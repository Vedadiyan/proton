package options

import (
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var selectSelector *protoimpl.ExtensionInfo
var querySelector *protoimpl.ExtensionInfo
var notMappedSelector *protoimpl.ExtensionInfo
var bindToSelector *protoimpl.ExtensionInfo
var typeSelector *protoimpl.ExtensionInfo
var indexSelector *protoimpl.ExtensionInfo

type ProtonOptions struct {
	selectAttr      string
	reduceAttr      string
	notMapped       bool
	bindToAttr      string
	typeAttr        string
	indexAttr       string
	fieldDescriptor protoreflect.FieldDescriptor
}

func init() {
	selectSelector = &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51234,
		Name:          "select",
		Tag:           "bytes,51234,opt,name=jsonField",
	}
	querySelector = &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51235,
		Name:          "query",
		Tag:           "bytes,51235,opt,name=reduce",
	}
	notMappedSelector = &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51236,
		Name:          "not_mapped",
		Tag:           "bytes,51236,opt,name=not_mapped",
	}
	bindToSelector = &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51237,
		Name:          "bind_to",
		Tag:           "bytes,51237,opt,name=bind_to",
	}
	typeSelector = &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51238,
		Name:          "type",
		Tag:           "bytes,51238,opt,name=type",
	}
	indexSelector = &protoimpl.ExtensionInfo{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51239,
		Name:          "index",
		Tag:           "bytes,51238,opt,name=index",
	}
}

func New(selectorAttr string, reduceAttr string, notMapped bool, bindToAttr string, typeAttr string, indexAttr string, fd protoreflect.FieldDescriptor) *ProtonOptions {
	protonOptions := &ProtonOptions{
		selectAttr:      selectorAttr,
		reduceAttr:      reduceAttr,
		notMapped:       notMapped,
		bindToAttr:      bindToAttr,
		typeAttr:        typeAttr,
		indexAttr:       indexAttr,
		fieldDescriptor: fd,
	}
	return protonOptions
}

func GetOptions(fd protoreflect.FieldDescriptor) *ProtonOptions {
	options := fd.Options().(*descriptorpb.FieldOptions)
	selectAttr := proto.GetExtension(options, selectSelector).(string)
	reduceAttr := proto.GetExtension(options, querySelector).(string)
	notMappedSelector := proto.GetExtension(options, notMappedSelector).(string) == "true"
	bindToAttr := proto.GetExtension(options, bindToSelector).(string)
	typeAttr := proto.GetExtension(options, typeSelector).(string)
	indexAttr := proto.GetExtension(options, indexSelector).(string)
	return New(selectAttr, reduceAttr, notMappedSelector, bindToAttr, typeAttr, indexAttr, fd)
}

func (protonOptions ProtonOptions) HasSelectAttribute() bool {
	return protonOptions.selectAttr != ""
}

func (protonOptions ProtonOptions) HasReduceAttribute() bool {
	return protonOptions.reduceAttr != ""
}

func (protonOptions ProtonOptions) HasAttributes() bool {
	return protonOptions.HasReduceAttribute() || protonOptions.HasSelectAttribute()
}

func (protonOptions ProtonOptions) GetSelectAttribute() string {
	return protonOptions.selectAttr
}

func (protonOptions ProtonOptions) GetReduceAttribute() string {
	return protonOptions.reduceAttr
}
func (protonOptions ProtonOptions) GetFieldName() string {
	var fieldName string
	fieldName = protonOptions.fieldDescriptor.JSONName()
	if fieldName == "" {
		fieldName = protonOptions.fieldDescriptor.TextName()
	}
	return fieldName
}

func (protonOptions ProtonOptions) IsNotMapped() bool {
	return protonOptions.notMapped
}

func (protonOptions ProtonOptions) GetBindTo() string {
	return protonOptions.bindToAttr
}

func (protonOptions ProtonOptions) GetType() string {
	return protonOptions.typeAttr
}

func (protonOptions ProtonOptions) GetIndex(length int) (int, int, error) {
	str := strings.Split(protonOptions.indexAttr, ":")
	start := 0
	end := length
	switch len(str) {
	case 2:
		i, err := strconv.ParseInt(str[1], 10, 32)
		if err != nil {
			return start, end, err
		}
		value := int(i)
		if value < end {
			end = value
		}
		fallthrough
	case 1:
		{
			i, err := strconv.ParseInt(str[0], 10, 32)
			if err != nil {
				return start, end, err
			}
			value := int(i)
			if value < end {
				start = value
			} else {
				start = end
			}
		}
	}
	return start, end, nil
}
