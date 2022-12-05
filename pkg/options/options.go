package options

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var selectSelector *protoimpl.ExtensionInfo
var querySelector *protoimpl.ExtensionInfo
var notMappedSelector *protoimpl.ExtensionInfo

type ProtonOptions struct {
	selectAttr      string
	reduceAttr      string
	notMapped       bool
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
		Tag:           "bytes,51236,opt,name=reduce",
	}
}

func New(selectorAttr string, reduceAttr string, notMapped bool, fd protoreflect.FieldDescriptor) *ProtonOptions {
	protonOptions := &ProtonOptions{
		selectAttr:      selectorAttr,
		reduceAttr:      reduceAttr,
		notMapped:       notMapped,
		fieldDescriptor: fd,
	}
	return protonOptions
}

func GetOptions(fd protoreflect.FieldDescriptor) *ProtonOptions {
	options := fd.Options().(*descriptorpb.FieldOptions)
	selectAttr := proto.GetExtension(options, selectSelector).(string)
	reduceAttr := proto.GetExtension(options, querySelector).(string)
	notMappedSelector := proto.GetExtension(options, notMappedSelector).(string) == "true"
	return New(selectAttr, reduceAttr, notMappedSelector, fd)
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
