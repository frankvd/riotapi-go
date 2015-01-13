package riotapi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewApi(t *testing.T) {
	api := NewApi("apikey")
	if api == nil {
		t.Error("NewApi() returned nil")
	}
}

func TestClientCall(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.RequestURI())
	}))
	client := NewClient()
	client.BaseUrl = server.URL
	client.Endpoints["test"] = "/{param}"

	resp := client.Call("test", "test_param")

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if buf.String() != "/test_param?api_key=apikey" {
		t.Error("Call returned invalid response: ", buf.String())
	}

	client.BaseUrl = "sadkfljrlktjarkjt"
	defer func() {
		if recover() == nil {
			t.Error("Call should panic on http err")
		}
	}()
	client.Call("test")

}
