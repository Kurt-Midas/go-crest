package crest

import (
	"testing"
)

//Alliances
func Test_Alliances(t *testing.T) {
	var bodyA = `{ "totalCount_str" : "4", "pageCount" : 2, "items" : [ { "id_str" : "498125261", "shortName" : "TEST", "href" : "https://crest-tq.eveonline.com/alliances/498125261/", "id" : 498125261, "name" : "Test Alliance Please Ignore" }, { "id_str" : "498125211", "shortName" : "WAT", "href" : "https://crest-tq.eveonline.com/alliances/498125211/", "id" : 498125211, "name" : "Doesnt Matter Had Fun" } ], "next" : { "href" : "https://crest-tq.eveonline.com/alliances/?page=2" }, "totalCount" : 4, "pageCount_str" : "2" }`
	var bodyB = `{ "totalCount_str" : "4", "pageCount" : 2, "items" : [ { "id_str" : "498125212", "shortName" : "DCA", "href" : "https://crest-tq.eveonline.com/alliances/498125212/", "id" : 498125212, "name" : "Duntcare A" }, { "id_str" : "498125211", "shortName" : "DCB", "href" : "https://crest-tq.eveonline.com/alliances/498125213/", "id" : 498125213, "name" : "Duntcare B" } ], "totalCount" : 4, "pageCount_str" : "2" }`
	forceFlippableProxyResponse(bodyA, bodyB)

	itemlist, err := Alliances_All()
	if err != nil {
		t.Errorf("Test_Alliances :: FAIL :: err was '%v'\n", err)
	} else if len(itemlist) != 4 {
		t.Errorf("Test_Alliances :: FAIL :: length was %v, expected 4\n", len(itemlist))
	} else if itemlist[0].ID != 498125261 {
		t.Errorf("Test_Alliances :: FAIL :: some parameters may be bad, dumping : %+v\n", itemlist[0])
	} else {
		t.Log("Test_Alliances :: Success")
	}
}

//Alliance
func Test_Alliance(t *testing.T) {
	var body = `{ "startDate" : "2010-11-08T03:34:00", "corporationsCount" : 3, "description" : "An elite pvp alliance, come get some boys", "executorCorporation" : { "name" : "Covenant of the Phoenix", "isNPC" : false, "href" : "https://crest-tq.eveonline.com/corporations/98001096/", "id_str" : "98001096", "logo" : { "32x32" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_32.png" }, "64x64" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_64.png" }, "128x128" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_128.png" }, "256x256" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_256.png" } }, "id" : 98001096 }, "corporationsCount_str" : "11", "deleted" : false, "creatorCorporation" : { "name" : "Covenant of the Phoenix", "isNPC" : false, "href" : "https://crest-tq.eveonline.com/corporations/98001096/", "id_str" : "98001096", "logo" : { "32x32" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_32.png" }, "64x64" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_64.png" }, "128x128" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_128.png" }, "256x256" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_256.png" } }, "id" : 98001096 }, "url" : "", "id_str" : "99000020", "creatorCharacter" : { "name" : "Jan Shan", "isNPC" : false, "href" : "https://crest-tq.eveonline.com/characters/681594178/", "capsuleer" : { "href" : "https://crest-tq.eveonline.com/characters/681594178/capsuleer/" }, "portrait" : { "32x32" : { "href" : "http://imageserver.eveonline.com/Character/681594178_32.jpg" }, "64x64" : { "href" : "http://imageserver.eveonline.com/Character/681594178_64.jpg" }, "128x128" : { "href" : "http://imageserver.eveonline.com/Character/681594178_128.jpg" }, "256x256" : { "href" : "http://imageserver.eveonline.com/Character/681594178_256.jpg" } }, "id" : 681594178, "id_str" : "681594178" }, "corporations" : [{ "name" : "Covenant of the Phoenix", "isNPC" : false, "href" : "https://crest-tq.eveonline.com/corporations/98001096/", "id_str" : "98001096", "logo" : { "32x32" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_32.png" }, "64x64" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_64.png" }, "128x128" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_128.png" }, "256x256" : { "href" : "http://imageserver.eveonline.com/Corporation/98001096_256.png" } }, "id" : 98001096 }, { "name" : "Covenant of the Raptor", "isNPC" : false, "href" : "https://crest-tq.eveonline.com/corporations/98001098/", "id_str" : "98001098", "logo" : { "32x32" : { "href" : "http://imageserver.eveonline.com/Corporation/98001098_32.png" }, "64x64" : { "href" : "http://imageserver.eveonline.com/Corporation/98001098_64.png" }, "128x128" : { "href" : "http://imageserver.eveonline.com/Corporation/98001098_128.png" }, "256x256" : { "href" : "http://imageserver.eveonline.com/Corporation/98001098_256.png" } }, "id" : 98001098 }, { "name" : "Covenant of the Parrot", "isNPC" : false, "href" : "https://crest-tq.eveonline.com/corporations/98001091/", "id_str" : "98001091", "logo" : { "32x32" : { "href" : "http://imageserver.eveonline.com/Corporation/98001091_32.png" }, "64x64" : { "href" : "http://imageserver.eveonline.com/Corporation/98001091_64.png" }, "128x128" : { "href" : "http://imageserver.eveonline.com/Corporation/98001091_128.png" }, "256x256" : { "href" : "http://imageserver.eveonline.com/Corporation/98001091_256.png" } }, "id" : 98001091 } ], "shortName" : "COTP", "id" : 99000020, "name" : "Covenant of the Phoenix Alliance" }`
	forceProxyResponse(body)
	item, err := Alliance(98001096) //doesn't have to be accurate
	if err != nil {
		t.Errorf("Test_Alliance :: FAIL :: err was '%v'\n", err)
	} else if item.CorporationsCount != 3 || len(item.Corporations) != 3 || item.ExecutorCorporation.ID != 98001096 {
		t.Errorf("Test_Alliance :: FAIL :: item did not populate properly, dumping : %+v\n", item)
	} else {
		t.Log("Test_Alliance (not that one) :: Success")
	}
}

//DogmaAttributes

func Test_DogmaAttributes(t *testing.T) {
	var bodyA = `{ "totalCount_str": "4", "pageCount": 2, "items": [ { "id_str": "2", "href": "https://crest-tq.eveonline.com/dogma/attributes/2/", "id": 2, "name": "isOnline" }, { "id_str": "3", "href": "https://crest-tq.eveonline.com/dogma/attributes/3/", "id": 3, "name": "isOnline" } ], "next": { "href": "https://crest-tq.eveonline.com/dogma/attributes/?page=2" }, "totalCount": 4, "pageCount_str": "2" }`
	var bodyB = `{ "totalCount_str": "4", "pageCount": 2, "items": [ { "id_str": "5", "href": "https://crest-tq.eveonline.com/dogma/attributes/5/", "id": 5, "name": "isOnline" }, { "id_str": "6", "href": "https://crest-tq.eveonline.com/dogma/attributes/6/", "id": 6, "name": "isOnline" } ], "totalCount": 4, "pageCount_str": "2" }`
	forceFlippableProxyResponse(bodyA, bodyB)
	itemlist, err := DogmaAttributes_All()
	if err != nil {
		t.Errorf("Test_DogmaAttributes :: FAIL :: err was '%v'\n", err)
	} else if len(itemlist) != 4 {
		t.Errorf("Test_DogmaAttributes :: FAIL :: length was %v, expected 4\n", len(itemlist))
	} else if itemlist[0].ID != 2 || itemlist[1].ID != 3 {
		t.Errorf("Test_DogmaAttributes :: FAIL :: itemlist did not populate properly, dumping : %+v\n", itemlist)
	} else {
		t.Log("Test_DogmaAttributes :: Success")
	}
}

//DogmaAttribute

func Test_DogmaAttribute(t *testing.T) {
	var body = `{ "description": "Factor by which topspeed increases.", "unitID_str": "124", "displayName": "Maximum Velocity Bonus", "name": "speedFactor", "iconID": 1389, "unitID": 124, "iconID_str": "1389", "id": 20, "id_str": "20", "highIsGood": true, "stackable": false, "defaultValue": 1, "published": true }`
	forceProxyResponse(body)
	item, err := DogmaAttribute(4) //doesn't have to be accurate
	if err != nil {
		t.Errorf("Test_DogmaAttribute :: FAIL :: err was '%v'\n", err)
	} else if item.DisplayName != "Maximum Velocity Bonus" || item.ID != 20 {
		t.Errorf("Test_DogmaAttribute :: FAIL :: item did not populate properly, dumping : %+v\n", item)
	} else {
		t.Log("Test_DogmaAttribute :: Success")
	}
}

//DogmaEffects

func Test_DogmaEffects(t *testing.T) {
	var bodyA = `{ "totalCount_str": "4", "pageCount": 2, "items": [ { "id_str": "2", "href": "https://crest-tq.eveonline.com/dogma/attributes/2/", "id": 2, "name": "isOnline" }, { "id_str": "3", "href": "https://crest-tq.eveonline.com/dogma/attributes/3/", "id": 3, "name": "isOnline" } ], "next": { "href": "https://crest-tq.eveonline.com/dogma/attributes/?page=2" }, "totalCount": 4, "pageCount_str": "2" }`
	var bodyB = `{ "totalCount_str": "4", "pageCount": 2, "items": [ { "id_str": "5", "href": "https://crest-tq.eveonline.com/dogma/attributes/5/", "id": 5, "name": "isOnline" }, { "id_str": "6", "href": "https://crest-tq.eveonline.com/dogma/attributes/6/", "id": 6, "name": "isOnline" } ], "totalCount": 4, "pageCount_str": "2" }`
	forceFlippableProxyResponse(bodyA, bodyB)
	itemlist, err := DogmaEffects_All()
	if err != nil {
		t.Errorf("Test_DogmaEffects :: FAIL :: err was '%v'\n", err)
	} else if len(itemlist) != 4 {
		t.Errorf("Test_DogmaEffects :: FAIL :: length was %v, expected 4\n", len(itemlist))
	} else if itemlist[0].ID != 2 || itemlist[1].ID != 3 {
		t.Errorf("Test_DogmaEffects :: FAIL :: itemlist did not populate properly, dumping : %+v\n", itemlist)
	} else {
		t.Log("Test_DogmaEffects :: Success")
	}
}

//DogmaEffect

func Test_DogmaEffect(t *testing.T) {
	var body = `{ "effectCategory_str": "0", "postExpression_str": "501", "isAssistance": false, "description": "", "isOffensive": false, "disallowAutoRepeat": false, "isWarpSafe": false, "preExpression_str": "500", "electronicChance": false, "rangeChance": false, "effectCategory": 0, "id_str": "61", "postExpression": 501, "published": false, "preExpression": 500, "displayName": "", "id": 61, "name": "agilityBonus" }`
	forceProxyResponse(body)
	item, err := DogmaEffect(4) //doesn't have to be accurate
	if err != nil {
		t.Errorf("Test_DogmaEffect :: FAIL :: err was '%v'\n", err)
	} else if item.Name != "agilityBonus" || item.ID != 61 {
		t.Errorf("Test_DogmaEffect :: FAIL :: item did not populate properly, dumping : %+v\n", item)
	} else {
		t.Log("Test_DogmaEffect :: Success")
	}
}

//Incursions

func Test_Incursions(t *testing.T) {
	var bodyA = `{ "totalCount_str": "4", "items": [ { "aggressorFactionID": { "id_str": "500019", "href": "https://crest-tq.eveonline.com/alliances/500019/", "id": 500019, "name": "Sansha's Nation" }, "stagingSolarSystem": { "id_str": "30004434", "href": "https://crest-tq.eveonline.com/solarsystems/30004434/", "id": 30004434, "name": "C-0ND2" }, "influence": 0, "incursionType": "Incursion", "state": "Withdrawing", "hasBoss": false, "infestedSolarSystems": [ { "id_str": "30004425", "href": "https://crest-tq.eveonline.com/solarsystems/30004425/", "id": 30004425, "name": "Z-PNIA" }, { "id_str": "30004426", "href": "https://crest-tq.eveonline.com/solarsystems/30004426/", "id": 30004426, "name": "OCU4-R" } ], "constellation": { "id_str": "20000647", "href": "https://crest-tq.eveonline.com/constellations/20000647/", "id": 20000647, "name": "C45-9Y" } }, { "aggressorFactionID": { "id_str": "500019", "href": "https://crest-tq.eveonline.com/alliances/500019/", "id": 500019, "name": "Sansha's Nation" }, "stagingSolarSystem": { "id_str": "30004411", "href": "https://crest-tq.eveonline.com/solarsystems/30004411/", "id": 30004411, "name": "101-IDK" }, "influence": 0, "incursionType": "Incursion", "state": "Withdrawing", "hasBoss": false, "infestedSolarSystems": [ { "id_str": "30004425", "href": "https://crest-tq.eveonline.com/solarsystems/30004425/", "id": 30004425, "name": "Z-PNIA" }, { "id_str": "30004426", "href": "https://crest-tq.eveonline.com/solarsystems/30004426/", "id": 30004426, "name": "OCU4-R" } ], "constellation": { "id_str": "20000647", "href": "https://crest-tq.eveonline.com/constellations/20000647/", "id": 20000647, "name": "C45-9Y" } } ], "next": { "href": "https://crest-tq.eveonline.com/incursions/?page=2" }, "pageCount": 2, "pageCount_str": "2", "totalCount": 4 }`
	var bodyB = `{ "totalCount_str": "4", "items": [ { "aggressorFactionID": { "id_str": "500019", "href": "https://crest-tq.eveonline.com/alliances/500019/", "id": 500019, "name": "Sansha's Nation" }, "stagingSolarSystem": { "id_str": "30004434", "href": "https://crest-tq.eveonline.com/solarsystems/30004434/", "id": 30004434, "name": "C-0ND2" }, "influence": 0, "incursionType": "Incursion", "state": "Withdrawing", "hasBoss": false, "infestedSolarSystems": [ { "id_str": "30004425", "href": "https://crest-tq.eveonline.com/solarsystems/30004425/", "id": 30004425, "name": "Z-PNIA" }, { "id_str": "30004426", "href": "https://crest-tq.eveonline.com/solarsystems/30004426/", "id": 30004426, "name": "OCU4-R" } ], "constellation": { "id_str": "20000647", "href": "https://crest-tq.eveonline.com/constellations/20000647/", "id": 20000647, "name": "C45-9Y" } }, { "aggressorFactionID": { "id_str": "500019", "href": "https://crest-tq.eveonline.com/alliances/500019/", "id": 500019, "name": "Sansha's Nation" }, "stagingSolarSystem": { "id_str": "30004411", "href": "https://crest-tq.eveonline.com/solarsystems/30004411/", "id": 30004411, "name": "101-IDK" }, "influence": 0, "incursionType": "Incursion", "state": "Withdrawing", "hasBoss": false, "infestedSolarSystems": [ { "id_str": "30004425", "href": "https://crest-tq.eveonline.com/solarsystems/30004425/", "id": 30004425, "name": "Z-PNIA" }, { "id_str": "30004426", "href": "https://crest-tq.eveonline.com/solarsystems/30004426/", "id": 30004426, "name": "OCU4-R" } ], "constellation": { "id_str": "20000647", "href": "https://crest-tq.eveonline.com/constellations/20000647/", "id": 20000647, "name": "C45-9Y" } } ], "pageCount": 2, "pageCount_str": "2", "totalCount": 4 }`
	forceFlippableProxyResponse(bodyA, bodyB)
	itemlist, err := Incursions_All()
	if err != nil {
		t.Errorf("Test_Incursions :: FAIL :: err was '%v'\n", err)
	} else if len(itemlist) != 4 {
		t.Errorf("Test_Incursions :: FAIL :: length was %v, expected 4\n", len(itemlist))
	} else if len(itemlist[0].InfestedSolarSystems) != 2 || itemlist[0].State != "Withdrawing" {
		t.Errorf("Test_Incursions :: FAIL :: itemlist did not populate properly, dumping : %+v\n", itemlist)
	} else {
		t.Log("Test_Incursions :: Success")
	}
}

//industry facilities ?
//InsurancePrices

func Test_InsurancePrices(t *testing.T) {
	var bodyA = `{ "totalCount_str": "4", "items": [ { "type": { "id_str": "582", "href": "https://crest-tq.eveonline.com/types/582/", "id": 582, "name": "Bantam" }, "insurance": [ { "payout": 209809.5, "cost": 20980.95, "level": "Basic" }, { "payout": 251771.4, "cost": 41961.9, "level": "Standard" }, { "payout": 293733.3, "cost": 62942.85, "level": "Bronze" }, { "payout": 335695.2, "cost": 83923.8, "level": "Silver" }, { "payout": 377657.10000000003, "cost": 104904.75, "level": "Gold" }, { "payout": 419619, "cost": 125885.7, "level": "Platinum" } ] }, { "type": { "id_str": "12005", "href": "https://crest-tq.eveonline.com/types/12005/", "id": 12005, "name": "Ishtar" }, "insurance": [ { "payout": 2009809.5, "cost": 200980.95, "level": "Basic" }, { "payout": 2051771.4, "cost": 401961.9, "level": "Standard" }, { "payout": 2093733.3, "cost": 602942.85, "level": "Bronze" }, { "payout": 3035695.2, "cost": 803923.8, "level": "Silver" }, { "payout": 3077657.10000000003, "cost": 1004904.75, "level": "Gold" }, { "payout": 4019619, "cost": 1025885.7, "level": "Platinum" } ] } ], "next" : { "href" : "https://crest-tq.eveonline.com/alliances/?page=2" }, "pageCount": 2, "pageCount_str": "2", "totalCount": 4 }`
	var bodyB = `{ "totalCount_str": "4", "items": [ { "type": { "id_str": "582", "href": "https://crest-tq.eveonline.com/types/582/", "id": 582, "name": "Bantam" }, "insurance": [ { "payout": 209809.5, "cost": 20980.95, "level": "Basic" }, { "payout": 251771.4, "cost": 41961.9, "level": "Standard" }, { "payout": 293733.3, "cost": 62942.85, "level": "Bronze" }, { "payout": 335695.2, "cost": 83923.8, "level": "Silver" }, { "payout": 377657.10000000003, "cost": 104904.75, "level": "Gold" }, { "payout": 419619, "cost": 125885.7, "level": "Platinum" } ] }, { "type": { "id_str": "12005", "href": "https://crest-tq.eveonline.com/types/12005/", "id": 12005, "name": "Ishtar" }, "insurance": [ { "payout": 2009809.5, "cost": 200980.95, "level": "Basic" }, { "payout": 2051771.4, "cost": 401961.9, "level": "Standard" }, { "payout": 2093733.3, "cost": 602942.85, "level": "Bronze" }, { "payout": 3035695.2, "cost": 803923.8, "level": "Silver" }, { "payout": 3077657.10000000003, "cost": 1004904.75, "level": "Gold" }, { "payout": 4019619, "cost": 1025885.7, "level": "Platinum" } ] } ], "pageCount": 2, "pageCount_str": "2", "totalCount": 4 }`
	forceFlippableProxyResponse(bodyA, bodyB)
	itemlist, err := InsurancePrices_All()
	if err != nil {
		t.Errorf("Test_InsurancePrices :: FAIL :: err was '%v'\n", err)
	} else if len(itemlist) != 4 {
		t.Errorf("Test_InsurancePrices :: FAIL :: length was %v, expected 4\n", len(itemlist))
	} else if itemlist[0].Type.ID != 582 || itemlist[1].Type.ID != 12005 || len(itemlist[0].Insurance) != 6 || itemlist[0].Insurance[0].Level != "Basic" {
		t.Errorf("Test_InsurancePrices :: FAIL :: itemlist did not populate properly, dumping : %+v\n", itemlist)
	} else {
		t.Log("Test_InsurancePrices :: Success")
	}
}

//Opportunities

//SovStructures
//SovCampaigns
//Types
//Type
