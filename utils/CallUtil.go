package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Any interface{}

/*type Pageable struct {
	PageCount      int           `json:"pageCount"`
	PageCount_str  string        `json:"pageCount_str"`
	TotalCount     int           `json:"totalCount"`
	TotalCount_str string        `json:"totalCount_str"`
	Items          []interface{} `json:"items"`
	Next           SubHref       `json:"next"`
}*/

type Pageable interface {
	GetCurrentPageNumber() int
	GetTotalPageCount() int
	GetNext() SubHref
	GetItems() Any
	SetItems(Any)
}

type SubHref struct {
	Href string `json:"href"`
}

const CrestDomain string = "https://crest-tq.eveonline.com"

/**
 * [Builds a CREST *http.Request object. Checks for errors, parses stuff, adds auth and user agent headers]
 * @param {[type]} method      string  [description]
 * @param {[type]} path        string  [description]
 * @param {[type]} accesstoken string) (error,       *http.Request [description]
 */
func BuildCrestCallRequest(method string, rawurl string, accesstoken string, data url.Values) (error, *http.Request) {
	// var crestDomain = "https://crest-tq.eveonline.com"
	// var rawurl = crestDomain + path
	urlStr, urlErr := url.Parse(rawurl)
	if urlErr != nil {
		return urlErr, nil
	}
	req, reqErr := http.NewRequest(method, urlStr.String(), bytes.NewBufferString(data.Encode()))
	if reqErr != nil {
		return reqErr, nil
	}
	req.Header.Add("Authorization", "Bearer "+accesstoken)
	req.Header.Add("User-Agent", "github.com/kurt-midas/go-crest")
	return nil, req
}

/**
 * [RemoteCall description]
 * @param {[type]} request   *http.Request [description]
 * @param {[type]} returnObj Any)          (error,       interface{} [description]
 */
func RemoteCall(request *http.Request, returnObj Any) error {
	client := &http.Client{}
	resp, respErr := client.Do(request)
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
	fmt.Printf("Resp :: %v\n", respBytes)
	jsonErr := json.Unmarshal(respBytes, returnObj)
	fmt.Printf("Json :: %+v\n", returnObj)
	if jsonErr != nil {
		return jsonErr
	}
	return nil
}

/**
 * [RemoteCall description]
 * @param {[type]} request   *http.Request [description]
 * @param {[type]} returnObj Any)          (error,       interface{} [description]
 */
func RemotePageableCall(request *http.Request, pageObj Pageable) error {
	client := &http.Client{}

	// pageObj.Next = SubHref{request.URL.String()}

	// var nextUrl url.URL = request.URL.String()
	var anySlice = make([]Any, 0)
	for request.URL != nil && request.URL.String() != "" {
		resp, respErr := client.Do(request)
		if respErr != nil {
			return respErr
		}
		defer resp.Body.Close()
		//TODO: handle status
		fmt.Printf("Response status : %v\n", resp.StatusCode)
		respBytes, bytesErr := ioutil.ReadAll(resp.Body)
		if bytesErr != nil {
			return bytesErr
		}
		fmt.Printf("Resp :: %v\n", respBytes)
		jsonErr := json.Unmarshal(respBytes, pageObj)
		fmt.Printf("Json :: %+v\n", pageObj)
		if jsonErr != nil {
			return jsonErr
		}
		anySlice = append(anySlice, pageObj.GetItems())
		if pageObj.GetNext().Href == "" {
			request.URL = nil
		} else {
			request.URL, _ = url.Parse(pageObj.GetNext().Href)
		}
	}
	// pageObj.Items = anySlice
	pageObj.SetItems(anySlice)
	return nil
}
