package crest

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	// "net/http/httptest"
	// "net/url"
	"strings"
	"testing"
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

func Test_charContacts(t *testing.T) {
	mockClient := http.DefaultClient
	var body string = `{ "totalCount_str": "1024", "pageCount": 6, "items": [ { "standing": 3.5, "character": { "name": "0men666", "corporation": { "name": "Far-East Corporation", "isNPC": false, "href": "https://api-sisi.testeveonline.com/corporations/98196279/", "id_str": "98196279", "logo": { "32x32": { "href": "https://image.testeveonline.com/Corporation/98196279_32.png" }, "64x64": { "href": "https://image.testeveonline.com/Corporation/98196279_64.png" }, "128x128": { "href": "https://image.testeveonline.com/Corporation/98196279_128.png" }, "256x256": { "href": "https://image.testeveonline.com/Corporation/98196279_256.png" } }, "id": 98196279 }, "isNPC": false, "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "capsuleer": { "href": "https://api-sisi.testeveonline.com/characters/1600000294/capsuleer/" }, "portrait": { "32x32": { "href": "https://image.testeveonline.com/Character/1600000294_32.jpg" }, "64x64": { "href": "https://image.testeveonline.com/Character/1600000294_64.jpg" }, "128x128": { "href": "https://image.testeveonline.com/Character/1600000294_128.jpg" }, "256x256": { "href": "https://image.testeveonline.com/Character/1600000294_256.jpg" } }, "id": 1600000294, "id_str": "1600000294" }, "contact": { "id_str": "1600000294", "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "name": "0men666", "id": 1600000294 }, "href": "https://api-sisi.testeveonline.com/characters/92095466/contacts/1600000294/", "contactType": "Character", "watched": true, "blocked": false } ], "totalCount": 1024, "pageCount_str": "6" }`
	mockClient.Transport = newMockTransport(body)
	utilClient = mockClient

	var accesstoken = "12345"
	var characterid = 67890
	items, err := Character_Contacts_Get(accesstoken, characterid)
	if err != nil {
		t.Errorf("Unexpected error with cause %+v\n", err)
	} else {
		t.Errorf("Success, response size is %v\n", len(items))
	}
}

/*func serverTestTools(code int, jsonBody string) (*httptest.Server, string, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, jsonBody)
	}))
	fmt.Printf("serverUrl :: %v \n", server.URL)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	fmt.Printf("Server URL is :: %v \n", server.URL)
	httpClient := &http.Client{Transport: transport}
	return server, server.URL, httpClient
}
*/
