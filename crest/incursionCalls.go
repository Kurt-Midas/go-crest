package crest

import (
	"net/url"
	"strconv"
)

//incursions
//Multiple is probably redundant. Currently this is only a single page. For consistency though...
func Incursions_All() ([]IncursionItem, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/incursions/"
	items := make([]IncursionItem, 0)
	for rawurl != "" {
		returnObj := new(IncursionsFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func Incursions_Page(pagenum int) ([]IncursionItem, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/incursions/?page=" + strconv.Itoa(pagenum)
	returnObj := new(IncursionsFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}
