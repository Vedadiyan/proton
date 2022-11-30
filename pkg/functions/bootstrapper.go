package functions

import (
	godyn "github.com/vedadiyan/godyn/pkg"
	"github.com/vedadiyan/proton/pkg/models"
)

var godynCtx godyn.Godyn[models.ProtonArg]

func init() {
	functions := make(map[string]godyn.Expression[models.ProtonArg])
	functions["IF"] = If
	functions["MIN"] = Min
	functions["MIN_OF"] = MinOf
	functions["MAX"] = Max
	functions["MAX_OF"] = MaxOf
	functions["FIRST"] = First
	functions["LAST"] = Last
	functions["WHERE"] = Where
	functions["SELECT"] = Select
	functions["FROM"] = From
	functions["REPLACE"] = Replace
	godynCtx = godyn.New(functions)
}

func GetFunctions() godyn.Godyn[models.ProtonArg] {
	return godynCtx
}
