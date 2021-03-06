package crest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const CrestDomain string = "https://crest-tq.eveonline.com"
const crestTypeUrl string = "?type=" + CrestDomain + "/inventory/types/"

var utilClient = &http.Client{}

func AvoidWarning() {
	fmt.Println("callutils.go, I exist")
}

func BuildAuthRequest(method string, rawurl string, accesstoken string, data url.Values) (*http.Request, error) {
	fmt.Printf("BuildAuthRequest :: rawurl is '%v'\n", rawurl)
	if validMethod, _ := regexp.MatchString("(GET|POST|PUT|DELETE|PATCH)", method); !validMethod {
		return nil, errors.New("Invalid method type '" + method + "'")
	}
	if validUrl, _ := regexp.MatchString("(http|https)://.*", rawurl); !validUrl {
		return nil, errors.New("Invalid url string '" + rawurl + "'")
	}

	urlStr, urlErr := url.Parse(rawurl)
	if urlErr != nil {
		return nil, urlErr
	}
	req, reqErr := http.NewRequest(method, urlStr.String(), bytes.NewBufferString(data.Encode()))
	if reqErr != nil {
		return nil, reqErr
	}
	if accesstoken != "" {
		req.Header.Add("Authorization", "Bearer "+accesstoken)
		// return nil, errors.New("BuildCrestCallRequest requires an access token")
	}
	req.Header.Add("User-Agent", "github.com/kurt-midas/go-crest")
	return req, nil
}

func RemoteCall(request *http.Request, returnObj Any) error {
	// client := &http.Client{}
	resp, respErr := utilClient.Do(request)
	if respErr != nil {
		return respErr
	}
	//TODO: handle status
	fmt.Printf("Response status : %v\n", resp.StatusCode)
	defer resp.Body.Close()
	respBytes, bytesErr := ioutil.ReadAll(resp.Body)
	if bytesErr != nil {
		return bytesErr
	}
	// fmt.Printf("RemoteCall :: Resp :: %v\n", respBytes)
	// fmt.Printf("RemoteCall :: Pre-Json :: %+v\n", returnObj)
	jsonErr := json.Unmarshal(respBytes, &returnObj)
	// fmt.Printf("RemoteCall :: Json :: %+v\n", returnObj)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

func RequestCall(method string, rawurl string, accesstoken string, data url.Values, returnObj Any) error {
	req, reqErr := BuildAuthRequest(method, rawurl, accesstoken, data)
	if reqErr != nil {
		return reqErr
	}
	return RemoteCall(req, returnObj)
}
