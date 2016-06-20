package crest

import (
	"fmt"
	"net/url"
	"strconv"
)

/******************************
******Char Contacts Calls******
******************************/
/*
	Contacts
		GET
		POST
	Contact
		GET
		PUT
		DELETE
*/
func Character_Contacts_Get(accesstoken string, characterid int) ([]CharContactItems, error) {
	var method = "GET"
	var rawurl string = CrestDomain + "/characters/" + strconv.Itoa(characterid) + "/contacts/"
	fmt.Printf("RawURL : '%v'\n", rawurl)
	items := make([]CharContactItems, 0)
	//can do this by hitting the call once and then for-ing on Next && pageNum and make-ing on totalValues
	//It's less clean than simply looping here
	for {
		req, reqErr := BuildAuthRequest(method, rawurl, accesstoken, url.Values{})
		var returnObj = new(CharContactsFormat)
		if reqErr != nil {
			return items, reqErr
		}
		callErr := RemoteCall(req, returnObj)
		if callErr != nil {
			return items, callErr
		}
		items = append(items, returnObj.Items...)
		if returnObj.Next.Href == "" {
			return items, nil
		}
		rawurl = returnObj.Next.Href //next page
	}
}

/******************************
******Char Fittings Calls******
******************************/
/*
	Fittings
		GET
		POST
	Fitting
		GET
		DELETE
*/
func Character_Fittings_Get(accesstoken string, characterid int) ([]CharFittingsItem, error) {
	var method = "GET"
	var rawurl string = CrestDomain + "/characters/" + strconv.Itoa(characterid) + "/fittings/"
	fmt.Printf("RawURL : '%v'\n", rawurl)
	items := make([]CharFittingsItem, 0)
	//can do this by hitting the call once and then for-ing on Next && pageNum and make-ing on totalValues
	//It's less clean than simply looping here
	for {
		req, reqErr := BuildAuthRequest(method, rawurl, accesstoken, url.Values{})
		var returnObj = new(CharFittingsFormat)
		if reqErr != nil {
			return items, reqErr
		}
		callErr := RemoteCall(req, returnObj)
		if callErr != nil {
			return items, callErr
		}
		items = append(items, returnObj.Items...)
		if returnObj.Next.Href == "" {
			return items, nil
		}
		rawurl = returnObj.Next.Href //next page
	}
}
