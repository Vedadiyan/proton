package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	unmarshaller "github.com/vedadiyan/proton/pkg/unmarshaller"
)

func TestAirline(t *testing.T) {
	var mapper map[string]any
	wd, err := os.Getwd()
	if err != nil {
		t.Error("cloud not get working directory")
		return
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/../resources/airline.json", wd))
	if err != nil {
		t.Error("could not read source file")
		return
	}
	err = json.Unmarshal(file, &mapper)
	if err != nil {
		t.Error("could not deserialize json")
		return
	}
	oofers := ArrayOfOffers{}
	err = unmarshaller.Unmarshall(mapper, &oofers)
	if err != nil {
		t.Error(err)
		return
	}
	len := len(oofers.GetOffers())
	if len == 0 {
		t.Error("bad unmarshaller")
		return
	}
	t.Logf("%d unmarshalled", len)
}
