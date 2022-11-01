package functions

import (
	"fmt"

	godyn "github.com/vedadiyan/godyn/pkg"
	"github.com/vedadiyan/proton/pkg/models"
	"github.com/vedadiyan/proton/pkg/sentinel"
)

func Where(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.ANY, godyn.STRING, godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	selector := args[1].(string)
	selector = fixString(selector)
	argValue := args[2]
	finalValue := fmt.Sprintf("%v", argValue)
	finalValue = fixString(finalValue)
	if key, ok := args[0].(string); ok {
		res := where(getKeys(key), selector, finalValue, data)
		return res, nil
	}
	if key, ok := args[0].(map[string]any); ok {
		res := whereKey([]string{}, finalValue, key)
		return res, nil
	}
	if data, ok := args[0].([]any); ok {
		res := whereAlt(selector, finalValue, data)
		return res, nil
	}
	return nil, nil
}

func Select(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.ANY, godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	if args[0] == nil {
		return nil, nil
	}
	var res any
	switch d := args[0].(type) {
	case string:
		{
			res, err := Select(data, []any{data.GetData(), d})
			if err != nil {
				return nil, err
			}
			return Select(data, []any{res, args[1]})
		}
	case map[string]any:
		{
			key := args[1].(string)
			key = fixString(key)
			res = _select(d, key)
		}
	case []any:
		{
			output := make([]any, 0)
			for _, value := range d {
				res, err := Select(data, []any{value, args[1]})
				if err != nil {
					return nil, err
				}
				output = append(output, res)
			}
			res = output

		}
	}
	return res, nil
}

func From(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	k := args[0].(string)
	key := getKeys(k)
	return from(key, data)
}

func where(key []string, selector string, value string, data models.ProtonArg) []any {
	prev := make([]any, 0)
	reduce(key, data.GetData(), func(current *map[string]any) bool {
		if current != nil {
			if v, ok := (*current)[selector]; ok {
				if v == value {
					prev = append(prev, *current)
				}
			}
		}
		return true
	})
	return prev
}

func whereKey(key []string, value string, data models.ProtonArg) any {
	var prev any
	reduce(key, data, func(current *map[string]any) bool {
		if current != nil {
			if v, ok := (*current)[value]; ok {
				prev = v
				return false
			}
		}
		return true
	})
	return prev
}

func whereAlt(selector string, value string, data []any) []any {
	prev := make([]any, 0)
	for _, d := range data {
		if _d, ok := d.(map[string]any); ok {
			reduce([]string{selector}, _d, func(current *any) bool {
				if current != nil {
					if *current == value {
						prev = append(prev, _d)
					}
				}
				return true
			})
		}
	}
	return prev
}

func _select(data models.ProtonArg, key string) any {
	keys := getKeys(key)
	prev := make([]any, 0)
	reduce(keys, data, func(current *any) bool {
		if current != nil {
			prev = append(prev, *current)
		}
		return true
	})
	if len(prev) == 0 {
		return nil
	}
	if len(prev) == 1 {
		return prev[0]
	}
	return prev
}

func from(key []string, data models.ProtonArg) (any, error) {
	var prev any
	var err error
	index := 0
	var d map[string]any
	if key[0] == "$ROOT" {
		key = key[1:]
		d = data.GetSelf()
	} else {
		d = data.GetData()
	}
	reduce(key, d, func(current *any) bool {
		if index > 0 {
			err = sentinel.InvalidDataType()
			return false
		}
		if current != nil {
			prev = *current
		}
		index++
		return true
	})
	return prev, err
}
