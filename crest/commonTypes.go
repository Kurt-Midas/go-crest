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

type Any interface{}
