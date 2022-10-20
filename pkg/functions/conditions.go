package functions

import godyn "github.com/vedadiyan/godyn/pkg"

func If(data map[string]any, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.BOOL, godyn.ANY, godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	condition := args[0].(bool)
	whenTrue := args[1]
	whenFalse := args[2]
	if condition {
		return whenTrue, nil
	}
	return whenFalse, nil
}
