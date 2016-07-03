package crest

import (
	"net/url"
	"strconv"
)

func Alliances_All() ([]ItemTypeDescription, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/alliances/"
	items := make([]ItemTypeDescription, 0)
	for rawurl != "" {
		returnObj := new(AllianceListFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func Alliances_Page(pagenum int) ([]ItemTypeDescription, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/alliances/?page=" + strconv.Itoa(pagenum)
	returnObj := new(AllianceListFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func Alliance(allianceId int) (AllianceDetails, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/alliances/" + strconv.Itoa(allianceId) + "/"
	returnObj := new(AllianceDetails)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	if err != nil {
		return *returnObj, err
	}
	return *returnObj, err
}
