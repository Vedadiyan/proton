package content

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
	file, err := os.ReadFile(fmt.Sprintf("%s/../resources/sample.json", wd))
	if err != nil {
		t.Error("could not read source file")
		return
	}
	err = json.Unmarshal(file, &mapper)
	if err != nil {
		t.Error("could not deserialize json")
		return
	}
	activities := Content{}
	err = unmarshaller.Unmarshall(mapper["content"].(map[string]any), &activities)
	if err != nil {
		t.Error(err)
		return
	}

}
