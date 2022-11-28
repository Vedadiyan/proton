package conversions

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func unableToCastError(ttype string) error {
	return fmt.Errorf(fmt.Sprintf("unable to cast to %s", ttype))
}

func Convert(kind protoreflect.Kind, target any) (any, error) {
	switch kind {
	case protoreflect.BoolKind:
		{
			return toBoolean(target)
		}
	case protoreflect.DoubleKind:
		{
			return toDouble(target)
		}
	case protoreflect.BytesKind:
		{
			return toByte(target)
		}
	case protoreflect.FloatKind:
		{
			return toFloat(target)
		}
	case protoreflect.Fixed32Kind:
		fallthrough
	case protoreflect.Int32Kind:
		{
			return toInt(target)
		}
	case protoreflect.Fixed64Kind:
		fallthrough
	case protoreflect.Int64Kind:
		{
			return toInt64(target)
		}
	case protoreflect.StringKind:
		{
			return fmt.Sprintf("%v", target), nil
		}
	case protoreflect.Uint32Kind:
		{
			return toUInt(target)
		}
	case protoreflect.Uint64Kind:
		{
			return toUInt64(target)
		}
	}
	return target, nil
}

func toBoolean(input any) (bool, error) {
	if value, ok := input.(bool); ok {
		return value, nil
	}
	if value, ok := input.(string); ok {
		return strconv.ParseBool(value)
	}
	return false, unableToCastError("boolean")
}

func toDouble(input any) (float64, error) {
	if value, ok := input.(float64); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return float64(value), nil
	}
	if value, ok := input.(int16); ok {
		return float64(value), nil
	}
	if value, ok := input.(int64); ok {
		return float64(value), nil
	}
	if value, ok := input.(int8); ok {
		return float64(value), nil
	}
	if value, ok := input.(float32); ok {
		return float64(value), nil
	}
	if value, ok := input.(float64); ok {
		return float64(value), nil
	}
	if value, ok := input.(uint); ok {
		return float64(value), nil
	}
	if value, ok := input.(uint16); ok {
		return float64(value), nil
	}
	if value, ok := input.(uint32); ok {
		return float64(value), nil
	}
	if value, ok := input.(uint64); ok {
		return float64(value), nil
	}
	if value, ok := input.(uint8); ok {
		return float64(value), nil
	}
	if value, ok := input.(string); ok {
		return strconv.ParseFloat(value, 64)
	}
	return 0, unableToCastError("float64")
}

func toByte(input any) (byte, error) {
	if value, ok := input.(byte); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return byte(value), nil
	}
	if value, ok := input.(int16); ok {
		return byte(value), nil
	}
	if value, ok := input.(int64); ok {
		return byte(value), nil
	}
	if value, ok := input.(int8); ok {
		return byte(value), nil
	}
	if value, ok := input.(float32); ok {
		return byte(value), nil
	}
	if value, ok := input.(float64); ok {
		return byte(value), nil
	}
	if value, ok := input.(uint); ok {
		return byte(value), nil
	}
	if value, ok := input.(uint16); ok {
		return byte(value), nil
	}
	if value, ok := input.(uint32); ok {
		return byte(value), nil
	}
	if value, ok := input.(uint64); ok {
		return byte(value), nil
	}
	if value, ok := input.(uint8); ok {
		return byte(value), nil
	}
	if value, ok := input.(string); ok {
		cnv, err := strconv.ParseInt(value, 10, 8)
		return byte(cnv), err
	}
	return 0, unableToCastError("byte")
}

func toFloat(input any) (float32, error) {
	if value, ok := input.(float32); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return float32(value), nil
	}
	if value, ok := input.(int16); ok {
		return float32(value), nil
	}
	if value, ok := input.(int64); ok {
		return float32(value), nil
	}
	if value, ok := input.(int8); ok {
		return float32(value), nil
	}
	if value, ok := input.(float32); ok {
		return float32(value), nil
	}
	if value, ok := input.(float64); ok {
		return float32(value), nil
	}
	if value, ok := input.(uint); ok {
		return float32(value), nil
	}
	if value, ok := input.(uint16); ok {
		return float32(value), nil
	}
	if value, ok := input.(uint32); ok {
		return float32(value), nil
	}
	if value, ok := input.(uint64); ok {
		return float32(value), nil
	}
	if value, ok := input.(uint8); ok {
		return float32(value), nil
	}
	if value, ok := input.(string); ok {
		cnv, err := strconv.ParseFloat(value, 32)
		return float32(cnv), err
	}
	return 0, unableToCastError("float32")
}

func toInt(input any) (int32, error) {
	if value, ok := input.(int32); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return int32(value), nil
	}
	if value, ok := input.(int16); ok {
		return int32(value), nil
	}
	if value, ok := input.(int64); ok {
		return int32(value), nil
	}
	if value, ok := input.(int8); ok {
		return int32(value), nil
	}
	if value, ok := input.(float32); ok {
		return int32(value), nil
	}
	if value, ok := input.(float64); ok {
		return int32(value), nil
	}
	if value, ok := input.(uint); ok {
		return int32(value), nil
	}
	if value, ok := input.(uint16); ok {
		return int32(value), nil
	}
	if value, ok := input.(uint32); ok {
		return int32(value), nil
	}
	if value, ok := input.(uint64); ok {
		return int32(value), nil
	}
	if value, ok := input.(uint8); ok {
		return int32(value), nil
	}
	if value, ok := input.(string); ok {
		cnv, err := strconv.ParseInt(value, 10, 32)
		return int32(cnv), err
	}
	return 0, unableToCastError("int32")
}

func toInt64(input any) (int64, error) {
	if value, ok := input.(int64); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return int64(value), nil
	}
	if value, ok := input.(int32); ok {
		return int64(value), nil
	}
	if value, ok := input.(int16); ok {
		return int64(value), nil
	}
	if value, ok := input.(int64); ok {
		return int64(value), nil
	}
	if value, ok := input.(int8); ok {
		return int64(value), nil
	}
	if value, ok := input.(float32); ok {
		return int64(value), nil
	}
	if value, ok := input.(float64); ok {
		return int64(value), nil
	}
	if value, ok := input.(uint); ok {
		return int64(value), nil
	}
	if value, ok := input.(uint16); ok {
		return int64(value), nil
	}
	if value, ok := input.(uint32); ok {
		return int64(value), nil
	}
	if value, ok := input.(uint64); ok {
		return int64(value), nil
	}
	if value, ok := input.(uint8); ok {
		return int64(value), nil
	}
	if value, ok := input.(string); ok {
		return strconv.ParseInt(value, 10, 32)
	}
	return 0, unableToCastError("int64")
}

func toUInt(input any) (uint32, error) {
	if value, ok := input.(uint32); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return uint32(value), nil
	}
	if value, ok := input.(int16); ok {
		return uint32(value), nil
	}
	if value, ok := input.(int64); ok {
		return uint32(value), nil
	}
	if value, ok := input.(int8); ok {
		return uint32(value), nil
	}
	if value, ok := input.(float32); ok {
		return uint32(value), nil
	}
	if value, ok := input.(float64); ok {
		return uint32(value), nil
	}
	if value, ok := input.(uint); ok {
		return uint32(value), nil
	}
	if value, ok := input.(uint16); ok {
		return uint32(value), nil
	}
	if value, ok := input.(uint32); ok {
		return uint32(value), nil
	}
	if value, ok := input.(uint64); ok {
		return uint32(value), nil
	}
	if value, ok := input.(uint8); ok {
		return uint32(value), nil
	}
	if value, ok := input.(string); ok {
		cnv, err := strconv.ParseUint(value, 10, 32)
		return uint32(cnv), err
	}
	return 0, unableToCastError("uint32")
}

func toUInt64(input any) (uint64, error) {
	if value, ok := input.(uint64); ok {
		return value, nil
	}
	if value, ok := input.(int); ok {
		return uint64(value), nil
	}
	if value, ok := input.(int16); ok {
		return uint64(value), nil
	}
	if value, ok := input.(int64); ok {
		return uint64(value), nil
	}
	if value, ok := input.(int8); ok {
		return uint64(value), nil
	}
	if value, ok := input.(float32); ok {
		return uint64(value), nil
	}
	if value, ok := input.(float64); ok {
		return uint64(value), nil
	}
	if value, ok := input.(uint); ok {
		return uint64(value), nil
	}
	if value, ok := input.(uint16); ok {
		return uint64(value), nil
	}
	if value, ok := input.(uint32); ok {
		return uint64(value), nil
	}
	if value, ok := input.(uint64); ok {
		return uint64(value), nil
	}
	if value, ok := input.(uint8); ok {
		return uint64(value), nil
	}
	if value, ok := input.(string); ok {
		return strconv.ParseUint(value, 10, 64)
	}
	return 0, unableToCastError("uint64")
}
