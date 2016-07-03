package crest

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type mockTransport struct{ Body string }

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")
	response.Body = ioutil.NopCloser(strings.NewReader(t.Body))
	return response, nil
}
func newMockTransport(body string) http.RoundTripper {
	return &mockTransport{body}
}

type mockTransport_flippable struct {
	BodyActive    string
	BodySecondary string
}

func (t *mockTransport_flippable) RoundTrip(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusOK,
	}
	response.Header.Set("Content-Type", "application/json")
	response.Body = ioutil.NopCloser(strings.NewReader(t.BodyActive))
	var temp = t.BodyActive
	t.BodyActive = t.BodySecondary
	t.BodySecondary = temp
	return response, nil
}
func newMockFlippableTransport(bodyA string, bodyB string) http.RoundTripper {
	return &mockTransport_flippable{bodyA, bodyB}
}

func forceProxyResponse(body string) {
	mockClient := http.DefaultClient
	mockClient.Transport = newMockTransport(body)
	utilClient = mockClient
}

func forceFlippableProxyResponse(bodyA string, bodyB string) {
	mockClient := http.DefaultClient
	mockClient.Transport = newMockFlippableTransport(bodyA, bodyB)
	utilClient = mockClient
}
