package crest

import (
	"net/url"
	"strconv"
)

/******************************
**********Fleet Calls**********
******************************/

func Fleet_Base_id(accesstoken string, fleetid int) (FleetBaseFormat, error) {
	var rawurl string = CrestDomain + "/fleet/" + strconv.Itoa(fleetid)
	return Fleet_Base(accesstoken, rawurl)
}

//fleets naturally come with URL defined, so I'll just make this easy
func Fleet_Base(accesstoken string, fleeturl string) (FleetBaseFormat, error) {
	var method = "GET"
	// fmt.Printf("RawURL : '%v'\n", fleeturl)
	fleetInfo := new(FleetBaseFormat)
	err := RequestCall(method, fleeturl, accesstoken, url.Values{}, fleetInfo)
	return *fleetInfo, err
}

/******************************
*********Member Calls**********
******************************/
func Fleet_Members_id(accesstoken string, fleetid int) ([]FleetMemberItems, error) {
	return Fleet_Members(accesstoken, CrestDomain+"/fleet/"+strconv.Itoa(fleetid)+"/members/")
}
func Fleet_Members(accesstoken string, fleetmembersurl string) ([]FleetMemberItems, error) {
	var method = "GET"
	items := make([]FleetMemberItems, 0)
	for fleetmembersurl != "" {
		returnObj := new(FleetMembersFormat)
		err := RequestCall(method, fleetmembersurl, accesstoken, url.Values{}, returnObj)
		if err != nil {
			return nil, err
		}
		items = append(items, returnObj.Items...)
		fleetmembersurl = returnObj.Next.Href
	}
	return items, nil
}

/******************************
*********Wing Calls************
******************************/
func Fleet_Wings_id(accesstoken string, fleetid int) ([]FleetWingItems, error) {
	return Fleet_Wings(accesstoken, CrestDomain+"/fleet/"+strconv.Itoa(fleetid)+"/wings/")
}
func Fleet_Wings(accesstoken string, fleetwingsurl string) ([]FleetWingItems, error) {
	var method = "GET"
	items := make([]FleetWingItems, 0)
	for fleetwingsurl != "" {
		returnObj := new(FleetWingsFormat)
		err := RequestCall(method, fleetwingsurl, accesstoken, url.Values{}, returnObj)
		if err != nil {
			return nil, err
		}
		items = append(items, returnObj.Items...)
		fleetwingsurl = returnObj.Next.Href
	}
	return items, nil
}
