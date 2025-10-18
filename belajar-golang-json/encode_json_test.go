package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data any) {
	bytes, error := json.Marshal(data)
	if error != nil {
		panic(error)
	}

	fmt.Println(string(bytes))
}

func TestEndode(t *testing.T) {
	logJson("ucup")
	logJson(1)
	logJson(true)
	logJson([]string{"ucup", "otong"})
}
