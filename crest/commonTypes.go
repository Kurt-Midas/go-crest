package crest

import ()

type ImgStruct struct {
	Size32  SubHref `json:"32x32"`
	Size64  SubHref `json:"64x64"`
	Size128 SubHref `json:"128x128"`
	Size256 SubHref `json:"256x256"`
}

type SubHref struct {
	Href string `json:"href"`
}

type ItemTypeDescription struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
}

type ItemIdDetails struct {
	ID_str string `json:"id_str"`
	Href   string `json:"href"`
	ID     int    `json:"id"`
}

type Pageable struct {
	TotalCount_str string  `json:"totalCount_str"`
	PageCount      int     `json:"pageCount"`
	PageCount_str  string  `json:"pageCount_str"`
	TotalCount     int     `json:"totalCount"`
	Next           SubHref `json:"next"`
}

type CorpDetails struct {
	Name   string    `json:"name"`
	IsNPC  bool      `json:"isNPC"`
	Href   string    `json:"href"`
	ID_str string    `json:"id_str"`
	ID     int       `json:"id"`
	Logo   ImgStruct `json:"logo"`
}

type CharDetails struct {
	Name string `json:"name"`
	// Character CorpDetails `json:"character"`
	IsNPC     bool      `json:"isNPC"`
	Href      string    `json:"href"`
	Capsuleer SubHref   `json:"capsuleer"`
	Portrait  ImgStruct `json:"portrait"`
	ID        int       `json:"id"`
	ID_str    string    `json:"id_str"`
}

type TypeListContainer struct {
	Pageable
	Items []ItemTypeDescription `json:"items"`
}

type Any interface{}
