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
	origin := args[0]
	source := fixString(args[1].(string))
	target := fixString(args[2].(string))
	switch t := origin.(type) {
	case string:
		{
			return strings.ReplaceAll(t, source, target), nil
		}
	case []any:
		{
			array := make([]any, 0)
			for _, item := range t {
				if str, ok := item.(string); !ok {
					array = append(array, strings.ReplaceAll(str, source, target))
				}
			}
			return array, nil
		}
	default:
		{
			return nil, fmt.Errorf("invalid type")
		}
	}
}
