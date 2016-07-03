package crest

import (
	"net/url"
	"strconv"
)

//opportunities

func OpportunitiesTasks_All() ([]OpportunitiesTaskItem, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/opportunities/tasks/"
	items := make([]OpportunitiesTaskItem, 0)
	for rawurl != "" {
		returnObj := new(OpportunitiesTaskFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func OpportunitiesTasks_Page(pagenum int) ([]OpportunitiesTaskItem, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/opportunities/tasks/?page=" + strconv.Itoa(pagenum)
	returnObj := new(OpportunitiesTaskFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func OpportunitiesTask(opportunityid int) (OpportunitiesTaskItem, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/opportunities/tasks/" + strconv.Itoa(opportunityid) + "/"
	returnObj := new(OpportunitiesTaskItem)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	if err != nil {
		return *returnObj, err
	}
	return *returnObj, err
}

func OpportunitiesGroups_All() ([]OpportunityGroup, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/opportunities/groups/"
	items := make([]OpportunityGroup, 0)
	for rawurl != "" {
		returnObj := new(OpportunityGroupFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

func OpportunitiesGroups_Page(pagenum int) ([]OpportunityGroup, int, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/opportunities/groups/?page=" + strconv.Itoa(pagenum)
	returnObj := new(OpportunityGroupFormat)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	return returnObj.Items, returnObj.PageCount, err
}

func OpportunitiesGroup(groupid int) (OpportunityGroup, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/opportunities/groups/" + strconv.Itoa(groupid) + "/"
	returnObj := new(OpportunityGroup)
	err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
	if err != nil {
		return *returnObj, err
	}
	return *returnObj, err
}
