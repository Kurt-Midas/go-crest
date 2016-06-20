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

type Any interface{}
