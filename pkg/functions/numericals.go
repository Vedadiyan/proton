package functions

import (
	"fmt"
	"math"
	"strconv"

	godyn "github.com/vedadiyan/godyn/pkg"
	"github.com/vedadiyan/proton/pkg/models"
)

func Min(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	key := args[0].(string)
	res := min(getKeys(key), data)
	return res, nil
}

func MinOf(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING, godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	key := args[0].(string)
	d := args[1].(map[string]any)
	res := min(getKeys(key), d)
	return res, nil
}

func Max(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	key := args[0].(string)
	res := max(getKeys(key), data)
	return res, nil
}

func MaxOf(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING, godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	key := args[0].(string)
	d := args[1].(map[string]any)
	res := max(getKeys(key), d)
	return res, nil
}

func Sum(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	key := args[0]
	switch value := key.(type) {
	case string:
		{
			res := sum(getKeys(value), data)
			return res, nil
		}
	case []float64:
		{
			var i float64 = 0
			for _, _value := range value {
				i += _value
			}
			return i, nil
		}
	default:
		{
			return nil, fmt.Errorf("invalid data type")
		}
	}
}

func ToNumber(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	key := args[0].(string)
	return toNumber(getKeys(key), data)
}

func min(key []string, data models.ProtonArg) float64 {
	var prev float64 = math.MaxInt64
	reduce(key, data.GetData(), func(current *float64) bool {
		if current != nil {
			if *current < prev {
				prev = *current
			}
		}
		return true
	})
	return prev
}

func max(key []string, data models.ProtonArg) float64 {
	var prev float64 = math.MaxInt64
	reduce(key, data.GetData(), func(current *float64) bool {
		if current != nil {
			if *current > prev {
				prev = *current
			}
		}
		return true
	})
	return prev
}

func sum(key []string, data models.ProtonArg) float64 {
	var prev float64 = 0
	reduce(key, data.GetData(), func(current *float64) bool {
		if current != nil {
			if *current > prev {
				prev += *current
			}
		}
		return true
	})
	return prev
}

func toNumber(key []string, data models.ProtonArg) ([]float64, error) {
	numbers := make([]float64, 0)
	var err error
	reduce(key, data.GetData(), func(current *string) bool {
		if current != nil {
			number, _err := strconv.ParseFloat(*current, 64)
			if _err != nil {
				err = _err
				return false
			}
			numbers = append(numbers, number)
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	return numbers, nil
}
