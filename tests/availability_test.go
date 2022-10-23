package test

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
	file, err := os.ReadFile(fmt.Sprintf("%s/resources/availability.json", wd))
	if err != nil {
		t.Error("could not read source file")
		return
	}
	err = json.Unmarshal(file, &mapper)
	if err != nil {
		t.Error("could not deserialize json")
		return
	}
	activities := ArrayOfAvailabilities{}
	err = unmarshaller.Unmarshall(mapper, &activities)
	if err != nil {
		t.Error(err)
		return
	}
	len := len(activities.GetAvailabilities())
	if len == 0 {
		t.Error("bad unmarshaller")
		return
	}
	t.Logf("%d unmarshalled", len)
	availability := activities.GetAvailabilities()[1]
	if availability.Id != "C BATLLO" {
		t.Error("invalid id")
	}
	if availability.Name != "Casa Batlló - Flexi Ticket" {
		t.Error("invalid name")
	}
	if availability.MinPrice != 33.33 {
		t.Error("invalid min price")
	}
	if availability.Thumbnail != "https://media.activitiesbank.com/11003/ENG/L/11003_5.jpg" {
		t.Error("invalid thumbnail")
	}
	if availability.Location.Longitude == 0 {
		t.Error("invalid location")
	}
	if availability.Location.Longitude == 0 {
		t.Error("invalid location")
	}
}
