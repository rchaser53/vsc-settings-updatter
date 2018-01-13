package main

import (
	"net/http"
	"testing"
)

type MockHTTPClient struct{}

func (m MockHTTPClient) Get(url string) (*http.Response, error) {
	return nil, NetworkError{Msg: "network error"}
}

func TestNetworkError(t *testing.T) {
	pullJSON := CreatePullSettingsJson(MockHTTPClient{})
	err := pullJSON("test.com", "destPath")

	switch err.(type) {
	case NetworkError:
		return
	default:
		t.Error("err should be NetworkError")
	}
}
