package crest

import (
	"net/url"
	"strconv"
)

//dogma attributes

func DogmaAttributes_All() ([]ItemTypeDescription, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/dogma/attributes/"
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

func DogmaAttributes_Paged(pagenum int) ([]ItemTypeDescription, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "dogma/attributes/?page=" + strconv.Itoa(pagenum)
	returnObj := new(TypeListContainer)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func DogmaAttribute(attributeid int) (DogmaAttributeFormat, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/dogma/attributes/" + strconv.Itoa(attributeid) + "/"
	returnObj := new(DogmaAttributeFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	if err != nil {
		return *returnObj, err
	}
	return *returnObj, err
}

//dogma effects

func DogmaEffects_All() ([]ItemTypeDescription, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/dogma/effects/"
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

func DogmaEffects_Page(pagenum int) ([]ItemTypeDescription, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/dogma/effects/?page=" + strconv.Itoa(pagenum)
	returnObj := new(TypeListContainer)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func DogmaEffect(effectid int) (DogmaEffectFormat, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/dogma/effects/" + strconv.Itoa(effectid) + "/"
	returnObj := new(DogmaEffectFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	if err != nil {
		return *returnObj, err
	}
	return *returnObj, err
}
