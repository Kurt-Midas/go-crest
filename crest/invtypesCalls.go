package crest

import (
	"net/url"
	"strconv"
)

//types

func InventoryTypes_All() ([]ItemTypeDescription, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/inventory/types/"
	items := make([]ItemTypeDescription, 0)
	for rawurl != "" {
		returnObj := new(TypeListContainer)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func InventoryTypes_Page(pagenum int) ([]ItemTypeDescription, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/inventory/types/?page=" + strconv.Itoa(pagenum)
	returnObj := new(TypeListContainer)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func InventoryType(typeid int) (InvTypeDetails, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/inventory/types/" + strconv.Itoa(typeid) + "/"
	returnObj := new(InvTypeDetails)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	if err != nil {
		return *returnObj, err
	}
	return *returnObj, err
}
