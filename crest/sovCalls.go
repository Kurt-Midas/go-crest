package crest

import (
	"net/url"
	"strconv"
)

//sovereignity

func SovereigntyStructures_All() ([]SovStructureItem, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/sovereignty/structures/"
	items := make([]SovStructureItem, 0)
	for rawurl != "" {
		returnObj := new(SovStructuresFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func SovereigntyStructures_Page(pagenum int) ([]SovStructureItem, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/sovereignty/structures/?page=" + strconv.Itoa(pagenum)
	returnObj := new(SovStructuresFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func SovereigntyCampaigns_All() ([]SovCampaignItem, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/sovereignty/campaigns/"
	items := make([]SovCampaignItem, 0)
	for rawurl != "" {
		returnObj := new(SovCampaignsFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func SovereigntyCampaigns_Page(pagenum int) ([]SovCampaignItem, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/sovereignty/campaigns/"
	returnObj := new(SovCampaignsFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}
