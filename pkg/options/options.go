package options

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var selectSelector *protoimpl.ExtensionInfo
var querySelector *protoimpl.ExtensionInfo

type ProtonOptions struct {
	selectAttr      string
	reduceAttr      string
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
		Tag:           "bytes,51234,opt,name=reduce",
	}
}

func New(selectorAttr string, reduceAttr string, fd protoreflect.FieldDescriptor) *ProtonOptions {
	protonOptions := &ProtonOptions{
		selectAttr:      selectorAttr,
		reduceAttr:      reduceAttr,
		fieldDescriptor: fd,
	}
	return protonOptions
}

func GetOptions(fd protoreflect.FieldDescriptor) *ProtonOptions {
	options := fd.Options().(*descriptorpb.FieldOptions)
	selectAttr := proto.GetExtension(options, selectSelector).(string)
	reduceAttr := proto.GetExtension(options, querySelector).(string)
	return New(selectAttr, reduceAttr, fd)
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
