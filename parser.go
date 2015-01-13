package riotapi

import (
	"encoding/json"
	"io"
)

// Parser interface
type Parser interface {
	Parse(response io.Reader, ret interface{})
}

// JSON parser
type JsonParser struct {
}

// Parse the json into the object
func (j *JsonParser) Parse(response io.Reader, ret interface{}) {
	json.NewDecoder(response).Decode(ret)
}
