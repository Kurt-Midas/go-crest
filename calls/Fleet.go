package calls

import (
	"fmt"
	"github.com/kurt-midas/go-crest/utils"
	"net/http"
	"net/url"
)

func AvoidWarning() {
	fmt.Printf("Fleet.go, package fleet, avoiding warning")
}

/**
 * 	Legend:
 * 		TODO
 * 		Calls
 * 		Response Structs
 */

/*
TODO:
	there should be an optional "options" thing for user agent and stuff
	is pageCount relevant? Do I need to handle that?

List of calls
	//Fleet Base GET
	Fleet Base PUT
	//Fleet Wings GET
	Fleet Wings POST //create new wing
	Fleet WingID DELETE
	Fleet WingID PUT //rename
	Fleet Wing Squad POST //create new squads
	Fleet Wing SquadID DELETE
	Fleet Wing SquadID PUT //rename
	//Fleet Members GET
	Fleet Members POST //invite
	Fleet MemberID PUT //move
	Fleet MemberID DELETE //kick
*/

/***************************
 ********Calls begin********
 ***************************/

func FleetBaseGetCall(rawfleetstr string, accesstoken string) (error, FleetBaseFormat) {
	urlStr, urlErr := url.Parse(rawfleetstr) //is this safe?
	if urlErr != nil {
		return urlErr, FleetBaseFormat{}
	}
	//TODO: overload to int method
	req, reqErr := http.NewRequest("GET", urlStr.String(), nil) //nil okay?
	req.Header.Add("Authorization", "Bearer "+accesstoken)
	req.Header.Add("User-Agent", "github.com/kurt-midas/go-crest")
	if reqErr != nil {
		return reqErr, FleetBaseFormat{}
	}
	var respStruct = new(FleetBaseFormat)
	callErr := utils.RemoteCall(req, *respStruct)
	if callErr != nil {
		return callErr, *respStruct
	}
	return nil, *respStruct
}

func FleetWingsGetCall(fleetid int, accesstoken string) (error, FleetWingsFormat) {
	reqErr, req := utils.BuildCrestCallRequest("GET", "/fleets/"+string(fleetid)+"/wings/", accesstoken, url.Values{})
	if reqErr != nil {
		return reqErr, FleetWingsFormat{}
	}
	var respStruct = new(FleetWingsFormat)
	callErr := utils.RemoteCall(req, *respStruct)
	if callErr != nil {
		return callErr, *respStruct
	}
	return nil, *respStruct
}

func FleetMemberGetCall(fleetid int, accesstoken string) (error, FleetMembersFormat) {
	reqErr, req := utils.BuildCrestCallRequest("GET", "/fleets/"+string(fleetid)+"/members/", accesstoken, url.Values{})
	if reqErr != nil {
		return reqErr, FleetMembersFormat{}
	}
	var respStruct = new(FleetMembersFormat)
	callErr := utils.RemoteCall(req, *respStruct)
	if callErr != nil {
		return callErr, *respStruct
	}
	return nil, *respStruct
}

/******************************
 *******Response Structs*******
 ******************************/

type FleetBaseFormat struct {
	IsVoiceEnabled bool          `json:"isVoiceEnabled"`
	Motd           string        `json:"motd"`
	IsFreeMove     bool          `json:"isFreeMove"`
	IsRegistered   bool          `json:"isRegistered"`
	Members        utils.SubHref `json:"members"`
	Wings          utils.SubHref `json:"wings"`
}

type FleetMembersFormat struct {
	TotalCount_str string             `json:"totalCount_str"`
	Items          []FleetMemberItems `json:"items"`
	PageCount      int                `json:"pageCount"`
	PageCount_str  string             `json:"pageCount_str"`
	TotalCount     int                `json:"totalCount"`
}

type FleetMemberItems struct {
	BoosterID      int                    `json:"boosterID"`
	BoosterID_str  string                 `json:"boosterID_str"`
	BoosterName    string                 `json:"boosterName"`
	Character      FleetMemberCharacter   `json:"character"`
	Href           string                 `json:"href"`
	JoinTime       string                 `json:"joinTime"`
	RoleID         int                    `json:"roleID"`
	RoleID_str     string                 `json:"roleID_str"`
	RoleName       string                 `json:"roleName"`
	Ship           FleetMemberShip        `json:"ship"`
	SolarSystem    FleetMemberSolarSystem `json:"solarSystem"`
	SquadID        int                    `json:"squadID"`
	SquadID_str    string                 `json:"squadID_str"`
	TakesFleetWarp bool                   `json:"takesFleetWarp"`
	WingID         int                    `json:"wingID"`
	WingID_str     string                 `json:"wingID_str"`
}

type FleetMemberSolarSystem struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}

type FleetMemberCharacter struct {
	Name      string        `json:"name"`
	IsNPC     bool          `json:"isNPC"`
	Href      string        `json:"href"`
	ID_str    string        `json:"id_str"`
	ID        int           `json:"id"`
	Capsuleer utils.SubHref `json:"capsuleer"`
}

type FleetMemberShip struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}

type FleetWingsFormat struct {
	TotalCount_str string           `json:"totalCount_str"`
	Items          []FleetWingItems `json:"items"`
	PageCount      int              `json:"pageCount"`
	PageCount_str  int              `json:"pageCount_str"`
	TotalCount     int              `json:"totalCount"`
}

type FleetWingItems struct {
	Name       string                  `json:"name"`
	Href       string                  `json:"href"`
	SquadsList []FleetWingSquadDetails `json:"squadsList"`
	ID_str     string                  `json:"id_str"`
	Squads     utils.SubHref           `json:"squads"`
	ID         int                     `json:"id"`
}

type FleetWingSquadDetails struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}
