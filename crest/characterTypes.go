package crest

import ()

/******************************
******Char Contacts types******
******************************/
type CharContactsFormat struct {
	Pageable
	Items []CharContactItems `json:"items"`
}
type CharContactItems struct {
	Standing    float32             `json:"standing"`
	Character   CharDetails         `json:"character"`
	Contact     ItemTypeDescription `json:"contact"`
	Href        string              `json:"href"`
	ContactType string              `json:"contactType"`
	Watched     bool                `json:"watched"`
	Blocked     bool                `json:"blocked"`
}

/******************************
******Char Fittings types******
******************************/
type CharFittingsFormat struct {
	Pageable
	Items []CharFittingsItem `json:"items"`
}
type CharFittingsItem struct {
	Description   string              `json:"description"`
	FittingID_str string              `json:"fittingID_str"`
	Items         []CharFittingModule `json:"items"`
	FittingID     int                 `json:"fittingID"`
	Ship          ItemTypeDescription `json:"ship"`
	Href          string              `json:"href"`
	Name          string              `json:"name"`
}

type CharFittingModule struct {
	Flag         int                 `json:"flag"`
	Flag_str     string              `json:"flag_str"`
	Quantity     int                 `json:"quantity"`
	Quantity_str string              `json:"quantity_str"`
	Type         ItemTypeDescription `json:"type"`
}

/******************************
******Char Location types******
******************************/

type CharLocationType struct {
	SolarSystem ItemTypeDescription `json:"solarSystem"`
}

/******************************
****Char Opportunities Types***
******************************/

type CharOpportunitiesFormat struct {
	Pageable
	Items []CharOpportunitiesItem `json:"items"`
}
type CharOpportunitiesItem struct {
	Completed string                `json:"completed"`
	Task      CharOpportunitiesTask `json:"task"`
}
type CharOpportunitiesTask struct {
	Href   string `json:"href"`
	ID     int    `json:"id"`
	ID_str string `json:"id_str"`
}
