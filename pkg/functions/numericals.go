package functions

import (
	"math"

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
