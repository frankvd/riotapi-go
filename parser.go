package riotapi

import (
	"encoding/json"
	"io"
)

type Parser interface {
	Parse(response string) interface{}
}

type JsonParser struct {
}

func (j *JsonParser) Parse(response io.ReadCloser, ret interface{}) {
	json.NewDecoder(response).Decode(ret)
}
