package functions

import (
	godyn "github.com/vedadiyan/godyn/pkg"
	"github.com/vedadiyan/proton/pkg/models"
)

func First(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	if key, ok := args[0].(string); ok {
		res := first(getKeys(key), data)
		return res, nil
	}
	if array, ok := args[0].([]any); ok {
		if len(array) == 0 {
			return nil, nil
		}
		return array[0], nil
	}
	return nil, nil
}

func Last(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	if key, ok := args[0].(string); ok {
		res := last(getKeys(key), data)
		return res, nil
	}
	if array, ok := args[0].([]any); ok {
		return array[len(array)-1], nil
	}
	return nil, nil
}

func first(key []string, data models.ProtonArg) any {
	var prev any
	reduce(key, data.GetData(), func(current *any) bool {
		if current != nil {
			prev = *current
		}
		return false
	})
	return prev
}

func last(key []string, data map[string]any) any {
	var prev any
	reduce(key, data["data"].(map[string]any), func(current *any) bool {
		if current != nil {
			prev = *current
		}
		return true
	})
	return prev
}
