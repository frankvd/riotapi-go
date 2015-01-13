package riotapi

import (
	"strings"
	"testing"
)

func TestJsonParser(t *testing.T) {
	parser := new(JsonParser)

	jsonstring := "{\"message\": \"PASSED\"}"
	reader := strings.NewReader(jsonstring)

	var val map[string]string

	parser.Parse(reader, &val)

	if val["message"] != "PASSED" {
		t.Error("Expected message PASSED got: ", val["message"])
	}
}
