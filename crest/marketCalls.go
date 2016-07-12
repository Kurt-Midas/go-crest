package crest

import (
	"net/url"
	"strconv"
)

//   ./market/types : gets all market items and some details (group, name, icon, id)
//	 ./market/types/<typeId:typeIdType>/ : details for one type
//	 ./market/prices/ : adjusted, average price for all market types
//	 ./market/groups/ : all market groups, href for types
//	 ./market/groups/<groupId:groupIdType>/ : one group
//	 ./market/<regionId:regionIdType>/orders/<"sell"|"buy">/<crest-type-url>/
//	 	takes a ?type=https://crest-tq.eveonline.com/inventory/types/??/, returns buy/sell orders
//	 ./market/<regionId>/orders/<all|crest-type-url>
//	 	returns ALL market orders
//	 ./market/<regionId:regionIdType>/history/<crest-type-url>
//	 	returns ordercount, low, high, avg, volume for 30 days

//just doing the important subset, ignoring string duplication, etc
type MarketOrderFormat struct {
	Pageable
	Items []MarketOrderType `json:"items"`
}

type MarketOrderType struct {
	Buy           bool    `json:"buy"`
	Issued        string  `json:"issued"`
	Price         float64 `json:"price"`
	VolumeEntered int     `json:"volumeEntered"`
	MinVolume     int     `json:"minVolume"`
	Volume        int     `json:"volume"`
	Range         string  `json:"range"`
	Href          string  `json:"href"`
	Location      struct {
		Href string `json:"location"`
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
	Duration int `json:"duration"`
	Type     struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Href string `json:"href"`
	} `json:"type"`
	Id int `json:"id"`
}

//Typed market orders
func MarketOrders_Type_All(region int, typeId int) ([]MarketOrderType, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/market/" + strconv.Itoa(region) + "/orders/" + crestTypeUrl + strconv.Itoa(typeId) + "/"
	items := make([]MarketOrderType, 0)
	for rawurl != "" {
		returnObj := new(MarketOrderFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

//typed history
func MarketHistory_Type(region int, typeId int) ([]MarketTypeHistory, error) {
	var method = "GET"
	var rawurl = CrestDomain + "/market/" + strconv.Itoa(region) + "/history/" + crestTypeUrl + strconv.Itoa(typeId) + "/"
	items := make([]MarketTypeHistory, 0)
	for rawurl != "" {
		returnObj := new(MarketHistoryFormat)
		err := RequestCall(method, rawurl, "", url.Values{}, returnObj)
		if err != nil {
			return items, err
		}
		items = append(items, returnObj.Items...)
		rawurl = returnObj.Next.Href
	}
	return items, nil
}

type MarketHistoryFormat struct {
	Pageable
	Items []kMarketTypeHistory `json:"items"`
}

type MarketTypeHistory struct {
	OrderCount int     `json:"orderCount"`
	LowPrice   float64 `json:"lowPrice"`
	HighPrice  float64 `json:"highPrice"`
	AvgPrice   float64 `json:"avgPrice"`
	Volume     int     `json:"volume"`
	Date       string  `json:"date"`
}
