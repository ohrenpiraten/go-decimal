package decimal

import (
	"encoding/json"
	"testing"
)

type TestDecHolder struct {
	Value Decimal
}

func TestMarshalJSON(t *testing.T) {

	data, err := json.Marshal(TestDecHolder{
		Value: New(12153, 3),
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "{\"Value\":\"12,153\"}" {
		t.Errorf("Fehler, json was %s", string(data))
	}
}

func TestUnMarshalJSON(t *testing.T) {

	var demo TestDecHolder
	err := json.Unmarshal([]byte("{\"Value\":\"12,153\"}"), &demo)
	if err != nil {
		t.Fatal(err)
	}
	if demo.Value.CompareTo(New(12153, 3)) != 0 {
		t.Errorf("Fehler, json was %s", demo.Value.Format())
	}
}
