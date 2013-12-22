package riotapi

import (
	"encoding/json"
)

type Parser interface {
	Parse(response string) interface{}
}

type JsonParser struct {
}

func (j *JsonParser) Parse(response string) {
	var json interface{}
	json.Unmarshal(response, &json)

	return json
}
