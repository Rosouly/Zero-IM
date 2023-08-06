package test

import (
	"encoding/json"
	"testing"
)

func TestJsonMessage(t *testing.T) {
	data := "hello world"

	bytes, err := json.Marshal(data)
	if err != nil {
		t.Error("json Marshal", err)
	}

	var data2 string
	json.Unmarshal(bytes, &data2)

	t.Log("data2", data2)
}
