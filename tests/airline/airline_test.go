package test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	marshaller "github.com/vedadiyan/proton/pkg/marshaller"
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
	offers := ArrayOfOffers{}
	err = unmarshaller.Unmarshall(mapper, &offers)
	if err != nil {
		t.Error(err)
		return
	}
	len := len(offers.GetOffers())
	if len == 0 {
		t.Error("bad unmarshaller")
		return
	}
	t.Logf("%d unmarshalled", len)
	d, err := marshaller.Marshall(&offers)
	if err != nil {
		t.Error("bad marshaller")
		return
	}
	_ = d
}
