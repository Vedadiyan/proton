package functions

import (
	"encoding/base64"
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
				if str, ok := item.(string); ok {
					array = append(array, strings.ReplaceAll(str, source, target))
				}
			}
			return array, nil
		}
	case nil:
		{
			return nil, nil
		}
	default:
		{
			return nil, fmt.Errorf("invalid type")
		}
	}
}

func Base64(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	origin := args[0]
	switch t := origin.(type) {
	case string:
		{
			return base64.URLEncoding.EncodeToString([]byte(t)), nil
		}
	case []any:
		{
			array := make([]any, 0)
			for _, item := range t {
				if str, ok := item.(string); ok {
					array = append(array, base64.URLEncoding.EncodeToString([]byte(str)))
				}
			}
			return array, nil
		}
	case nil:
		{
			return nil, nil
		}
	default:
		{
			return nil, fmt.Errorf("invalid type")
		}
	}
}

func Concat(data models.ProtonArg, args []any) (any, error) {
	err := godyn.ValidateArguments([]godyn.Type{godyn.STRING, godyn.ANY}, args)
	if err != nil {
		return nil, err
	}
	first := fixString(args[0].(string))
	second := args[1]
	switch t := second.(type) {
	case string:
		{
			return fmt.Sprintf("%s%s", first, t), nil
		}
	case []any:
		{
			array := make([]any, 0)
			for _, item := range t {
				if str, ok := item.(string); ok {
					array = append(array, fmt.Sprintf("%s%s", first, str))
				}
			}
			return array, nil
		}
	case nil:
		{
			return nil, nil
		}
	default:
		{
			return nil, fmt.Errorf("invalid type")
		}
	}
}
