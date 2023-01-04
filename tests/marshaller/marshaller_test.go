package marshaller

import (
	"testing"

	marshaller "github.com/vedadiyan/proton/pkg/marshaller"
)

func TestMarshaller(t *testing.T) {
	test := Test{}
	test.CheckIn = "123456"
	test.CheckOut = "454545"
	test.Latitude = 11111
	test.Latitude = 22222
	test.Adults = 10
	test.Children = []int32{50}
	var d map[string]any
	err := marshaller.Marshall(&test, &d)
	if err != nil {
		t.FailNow()
	}
}
