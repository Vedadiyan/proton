package sentinel

import "fmt"

func ObjectIsNotArray() error {
	return fmt.Errorf("object is not an array")
}

func TargetDoesNotExist() error {
	return fmt.Errorf("target does not exist")
}

func InvalidDataType() error {
	return fmt.Errorf("invalid data type")
}

func NotFound() error {
	return fmt.Errorf("not found")
}
