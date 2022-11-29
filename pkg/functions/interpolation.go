package functions

import (
	"fmt"
	"strings"

	godyn "github.com/vedadiyan/godyn/pkg"
	"github.com/vedadiyan/proton/pkg/models"
)

func Replace(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.ANY, godyn.STRING, godyn.STRING}, args)
	if err != nil {
		return nil, err
	}
	_source := args[0]
	source := fixString(args[1].(string))
	target := fixString(args[2].(string))
	_ = _source
	_ = source
	_ = target
	switch t := _source.(type) {
	case string:
		{
			return strings.ReplaceAll(t, source, target), nil
		}
	case []any:
		{
			if len(t) == 0 {
				return nil, fmt.Errorf("array is empty")
			}
			if _, ok := t[0].(string); !ok {
				return nil, fmt.Errorf("invalid array")
			}
			array := make([]any, 0)
			for _, item := range t {
				array = append(array, strings.ReplaceAll(item.(string), source, target))
			}
			return array, nil
		}
	default:
		{
			return nil, fmt.Errorf("invalid type")
		}
	}
}
