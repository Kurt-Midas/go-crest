package crest

import (
	"net/url"
	"strconv"
)

//insuranceprices

func InsurancePrices_All() ([]InsuranceInfo, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/insuranceprices/"
	items := make([]InsuranceInfo, 0)
	for rawurl != "" {
		returnObj := new(InsuranceInfoFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func InsurancePrices_Paged(pagenum int) ([]InsuranceInfo, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/insuranceprices/?page=" + strconv.Itoa(pagenum)
	returnObj := new(InsuranceInfoFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}
