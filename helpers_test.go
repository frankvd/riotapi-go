package riotapi

import (
	"io"
	"net/http"
)

type FakeClient struct {
	Response *http.Response
}

func (client *FakeClient) Call(endpoint string, params ...string) *http.Response {
	return client.Response
}

func (client *FakeClient) SetResponse(res *http.Response) {
	client.Response = res
}

type FakeParser struct {
	ParseFunc func(response io.Reader, ret interface{})
}

func (parser *FakeParser) Parse(response io.Reader, ret interface{}) {
	parser.ParseFunc(response, ret)
}

func NewFakeService() *Service {
	fakeClient := new(FakeClient)
	fakeClient.Response = new(http.Response)
	fakeParser := new(FakeParser)

	return &Service{Client: fakeClient, Parser: fakeParser}
}
