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
func Character_Contacts(accesstoken string, characterid int) ([]CharContactItems, error) {
	var method = "GET"
	var rawurl string = CrestDomain + "/characters/" + strconv.Itoa(characterid) + "/contacts/"
	fmt.Printf("RawURL : '%v'\n", rawurl)
	items := make([]CharContactItems, 0)
	//can do this by hitting the call once and then for-ing on Next && pageNum and make-ing on totalValues
	//It's less clean than simply looping here
	//I could use goroutines for lots of pages. No use case for now
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
		//GET
		POST - new fitting entirely
	Fitting
		GET - single fitting
		DELETE - deletes the resource
*/
func Character_Fittings(accesstoken string, characterid int) ([]CharFittingsItem, error) {
	var method = "GET"
	var rawurl string = CrestDomain + "/characters/" + strconv.Itoa(characterid) + "/fittings/"
	fmt.Printf("RawURL : '%v'\n", rawurl)
	items := make([]CharFittingsItem, 0)
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

//The example POST request wants everything, including href and name for types. Accurate?
//Fitting GET just returns a single one
func Character_Fitting(accesstoken string, characterid int, fittingid int) (CharFittingsItem, error) {
	var method = "GET"
	var rawurl string = CrestDomain + "/characters" + strconv.Itoa(characterid) +
		"/fittings/" + strconv.Itoa(fittingid) + "/"
	returnObj := new(CharFittingsItem)
	err := RequestCall(method, rawurl, accesstoken, url.Values{}, returnObj)
	return *returnObj, err
}

//Fitting DELETE works like GET but deletes a fitting
//This probably doesn't have a reply. Test it to be sure.
func Character_Fitting_Delete(accesstoken string, characterid int, fittingid int) (Any, error) {
	var method = "DELETE"
	var rawurl string = CrestDomain + "/characters" + strconv.Itoa(characterid) +
		"/fittings/" + strconv.Itoa(fittingid) + "/"
	returnObj := new(Any)
	err := RequestCall(method, rawurl, accesstoken, url.Values{}, returnObj)
	return returnObj, err
}

/******************************
******Char Location Calls******
******************************/
func Character_Location(accesstoken string, characterid int) (CharLocationType, error) {
	var method = "GET"
	var rawurl string = CrestDomain + "/characters/" + strconv.Itoa(characterid) + "/location/"
	// fmt.Printf("RawURL : '%v'\n", rawurl)
	solarSystemInfo := new(CharLocationType)
	err := RequestCall(method, rawurl, accesstoken, url.Values{}, solarSystemInfo)
	return *solarSystemInfo, err
}

/******************************
***Char Opportunities Calls****
******************************/
func Character_Opportunities(accesstoken string, characterid int) ([]CharOpportunitiesItem, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/characters/" + strconv.Itoa(characterid) + "/characterOpportunitiesRead/"
	items := make([]CharOpportunitiesItem, 0)
	for rawurl != "" {
		returnObj := new(CharOpportunitiesFormat)
		err := RequestCall(method, rawurl, accesstoken, url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

/******************************
******Char Waypoints Call******
******************************/
//POST only
// The example documentation only mentions a system destination.
// In-game it is possible to set stations as destination
// This merits further investigation
// /characters/<characterID:characterIdType>/navigation/waypoints/
