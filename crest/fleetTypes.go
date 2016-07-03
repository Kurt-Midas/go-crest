package crest

import ()

type FleetBaseFormat struct {
	IsVoiceEnabled bool    `json:"isVoiceEnabled"`
	Motd           string  `json:"motd"`
	IsFreeMove     bool    `json:"isFreeMove"`
	IsRegistered   bool    `json:"isRegistered"`
	Members        SubHref `json:"members"`
	Wings          SubHref `json:"wings"`
}

type FleetMembersFormat struct {
	Pageable
	Items []FleetMemberItems `json:"items"`
}

type FleetMemberItems struct {
	BoosterID      int                  `json:"boosterID"`
	BoosterID_str  string               `json:"boosterID_str"`
	BoosterName    string               `json:"boosterName"`
	Character      FleetMemberCharacter `json:"character"`
	Href           string               `json:"href"`
	JoinTime       string               `json:"joinTime"`
	RoleID         int                  `json:"roleID"`
	RoleID_str     string               `json:"roleID_str"`
	RoleName       string               `json:"roleName"`
	Ship           ItemTypeDescription  `json:"ship"`
	SolarSystem    ItemTypeDescription  `json:"solarSystem"`
	SquadID        int                  `json:"squadID"`
	SquadID_str    string               `json:"squadID_str"`
	TakesFleetWarp bool                 `json:"takesFleetWarp"`
	WingID         int                  `json:"wingID"`
	WingID_str     string               `json:"wingID_str"`
}

/*type FleetMemberSolarSystem struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}*/

type FleetMemberCharacter struct {
	Name      string  `json:"name"`
	IsNPC     bool    `json:"isNPC"`
	Href      string  `json:"href"`
	ID_str    string  `json:"id_str"`
	ID        int     `json:"id"`
	Capsuleer SubHref `json:"capsuleer"`
}

/*type FleetMemberShip struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}*/

/*type FleetWingsFormat struct {
	TotalCount_str string           `json:"totalCount_str"`
	Items          []FleetWingItems `json:"items"`
	PageCount      int              `json:"pageCount"`
	PageCount_str  int              `json:"pageCount_str"`
	TotalCount     int              `json:"totalCount"`
}*/

type FleetWingsFormat struct {
	Pageable
	Items []FleetWingItems `json:"items"`
}

type FleetWingItems struct {
	Name       string                `json:"name"`
	Href       string                `json:"href"`
	SquadsList []ItemTypeDescription `json:"squadsList"`
	ID_str     string                `json:"id_str"`
	Squads     SubHref               `json:"squads"`
	ID         int                   `json:"id"`
}

/*type FleetWingSquadDetails struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}*/
