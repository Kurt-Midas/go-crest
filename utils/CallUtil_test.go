package utils

import (
	"net/url"
	"testing"
)

func Test_BuildCrestCallRequest(t *testing.T) {
	cases := []struct {
		method      string
		rawurl      string
		accesstoken string
		data        url.Values
		isError     bool
	}{
		{"POST", CrestDomain + "/fleets/123456789/", "987654321", url.Values{}, false},
		{"", "", "", url.Values{}, true},                      //bad everything
		{"", CrestDomain + "/", "123", url.Values{}, true},    //bad method
		{"POST", CrestDomain + "", "123", url.Values{}, true}, //bad url
		{"POST", CrestDomain + "/", "", url.Values{}, true},   //no token
	}
	for i, c := range cases {
		err, req := BuildCrestCallRequest(c.method, c.rawurl, c.accesstoken, c.data)
		if !c.isError {
			if err != nil {
				t.Errorf("Unexpected error in case %v, error was %v", i, err)
			} else {
				if req.URL.String() == "" {
					t.Errorf("Failure :: success case, no error :: url was empty")
				}
				if req.Header.Get("Authorization") == "" {
					t.Errorf("Failure :: success case, no error :: 'Authorization' header not found")
				}
				if req.Header.Get("User-Agent") == "" {
					t.Errorf("Failure :: success case, no error :: 'User-Agent' header not found")
				}
			}
		} else if c.isError && err == nil {
			t.Errorf("Unexpected success in case %v, request was %v", i, req)
		}
	}
	// BuildCrestCallRequest(method, rawurl, accesstoken, data)
}
