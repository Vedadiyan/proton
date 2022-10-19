package helpers

import "google.golang.org/protobuf/reflect/protoreflect"

func FieldName(fd protoreflect.FieldDescriptor) string {
	fieldName := fd.JSONName()
	if fieldName == "" {
		fieldName = fd.TextName()
	}
	return fieldName
}
