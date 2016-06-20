package crest

import ()

/******************************
******Char Contacts types******
******************************/
type CharContactsFormat struct {
	PageCount      int                `json:"pageCount"`
	PageCount_str  string             `json:"pageCount_str"`
	TotalCount     int                `json:"totalCount"`
	TotalCount_str string             `json:"totalCount_str"`
	Items          []CharContactItems `json:"items"`
	Next           SubHref            `json:"next"`
}
type CharContactItems struct {
	Standing    float32             `json:"standing"`
	Character   CharContactItemChar `json:"character"`
	Contact     CharContactContact  `json:"contact"`
	Href        string              `json:"href"`
	ContactType string              `json:"contactType"`
	Watched     bool                `json:"watched"`
	Blocked     bool                `json:"blocked"`
}

type CharContactItemChar struct {
	Name      string          `json:"name"`
	Character CharContactCorp `json:"character"`
	IsNPC     bool            `json:"isNPC"`
	Href      string          `json:"href"`
	Capsuleer SubHref         `json:"capsuleer"`
	Portrait  ImgStruct       `json:"portrait"`
	ID        int             `json:"id"`
	ID_str    string          `json:"id_str"`
}

type CharContactCorp struct {
	Name   string    `json:"name"`
	IsNPC  bool      `json:"isNPC"`
	Href   string    `json:"href"`
	ID_str string    `json:"id_str"`
	ID     int       `json:"id"`
	Logo   ImgStruct `json:"logo"`
}

type CharContactContact struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	Name   string `json:"name"`
	ID     int    `json:"id"`
}

/******************************
******Char Fittings types******
******************************/

type CharFittingsFormat struct {
	PageCount      int                `json:"pageCount"`
	PageCount_str  string             `json:"pageCount_str"`
	TotalCount     int                `json:"totalCount"`
	TotalCount_str string             `json:"totalCount_str"`
	Items          []CharFittingsItem `json:"items"`
	Next           SubHref            `json:"next"`
}

type CharFittingsItem struct {
	Description   string              `json:"description"`
	FittingID_str string              `json:"fittingID_str"`
	Items         CharFittingModule   `json:"items"`
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
