package hotels

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	unmarshaller "github.com/vedadiyan/proton/pkg/unmarshaller"
)

func TestAvailability(t *testing.T) {
	var mapper map[string]any
	wd, err := os.Getwd()
	if err != nil {
		t.Error("cloud not get working directory")
		return
	}
	file, err := os.ReadFile(fmt.Sprintf("%s/../resources/hotels.json", wd))
	if err != nil {
		t.Error("could not read source file")
		return
	}
	err = json.Unmarshal(file, &mapper)
	if err != nil {
		t.Error("could not deserialize json")
		return
	}
	activities := ArrayOfResults{}
	err = unmarshaller.Unmarshall(mapper, &activities)
	if err != nil {
		t.Error(err)
		return
	}
}
