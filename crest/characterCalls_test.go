package crest

import (
	"testing"
)

/*************************
****Character Contacts****
*************************/

func Test_charContacts(t *testing.T) {
	var body string = `{ "totalCount_str": "1024", "pageCount": 6, "items": [ { "standing": 3.5, "character": { "name": "0men666", "corporation": { "name": "Far-East Corporation", "isNPC": false, "href": "https://api-sisi.testeveonline.com/corporations/98196279/", "id_str": "98196279", "logo": { "32x32": { "href": "https://image.testeveonline.com/Corporation/98196279_32.png" }, "64x64": { "href": "https://image.testeveonline.com/Corporation/98196279_64.png" }, "128x128": { "href": "https://image.testeveonline.com/Corporation/98196279_128.png" }, "256x256": { "href": "https://image.testeveonline.com/Corporation/98196279_256.png" } }, "id": 98196279 }, "isNPC": false, "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "capsuleer": { "href": "https://api-sisi.testeveonline.com/characters/1600000294/capsuleer/" }, "portrait": { "32x32": { "href": "https://image.testeveonline.com/Character/1600000294_32.jpg" }, "64x64": { "href": "https://image.testeveonline.com/Character/1600000294_64.jpg" }, "128x128": { "href": "https://image.testeveonline.com/Character/1600000294_128.jpg" }, "256x256": { "href": "https://image.testeveonline.com/Character/1600000294_256.jpg" } }, "id": 1600000294, "id_str": "1600000294" }, "contact": { "id_str": "1600000294", "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "name": "0men666", "id": 1600000294 }, "href": "https://api-sisi.testeveonline.com/characters/92095466/contacts/1600000294/", "contactType": "Character", "watched": true, "blocked": false } ], "totalCount": 1024, "pageCount_str": "6" }`
	forceProxyResponse(body)

	var accesstoken = "12345"
	var characterid = 67890
	items, err := Character_Contacts(accesstoken, characterid)
	if err != nil {
		t.Errorf("Test_charContacts :: FAIL :: Unexpected error with cause %+v\n", err)
	} else if len(items) != 1 {
		t.Errorf("Test_charContacts :: FAIL :: wanted length 1, got '%v'\n", len(items))
	} else if items[0].Standing != 3.5 {
		t.Errorf("Test_charContacts :: FAIL :: item 0 needed standing 3.5, got '%v'\n", items[0].Standing)
	} else {
		t.Logf("Test_charContacts :: Success, response size is %v\n", len(items))
	}
}

func Test_charContacts_paged(t *testing.T) {
	var bodyA string = `{ "totalCount_str": "1024", "pageCount": 6, "items": [ { "standing": 3.5, "character": { "name": "0men666", "corporation": { "name": "Far-East Corporation", "isNPC": false, "href": "https://api-sisi.testeveonline.com/corporations/98196279/", "id_str": "98196279", "logo": { "32x32": { "href": "https://image.testeveonline.com/Corporation/98196279_32.png" }, "64x64": { "href": "https://image.testeveonline.com/Corporation/98196279_64.png" }, "128x128": { "href": "https://image.testeveonline.com/Corporation/98196279_128.png" }, "256x256": { "href": "https://image.testeveonline.com/Corporation/98196279_256.png" } }, "id": 98196279 }, "isNPC": false, "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "capsuleer": { "href": "https://api-sisi.testeveonline.com/characters/1600000294/capsuleer/" }, "portrait": { "32x32": { "href": "https://image.testeveonline.com/Character/1600000294_32.jpg" }, "64x64": { "href": "https://image.testeveonline.com/Character/1600000294_64.jpg" }, "128x128": { "href": "https://image.testeveonline.com/Character/1600000294_128.jpg" }, "256x256": { "href": "https://image.testeveonline.com/Character/1600000294_256.jpg" } }, "id": 1600000294, "id_str": "1600000294" }, "contact": { "id_str": "1600000294", "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "name": "0men666", "id": 1600000294 }, "href": "https://api-sisi.testeveonline.com/characters/92095466/contacts/1600000294/", "contactType": "Character", "watched": true, "blocked": false } ], "totalCount": 1024, "pageCount_str": "6", "next":{ "href" : "https://crest-tq.eveonline.com/characters/67890/contacts/?page=2"} }`
	var bodyB string = `{ "totalCount_str": "1024", "pageCount": 6, "items": [ { "standing": 3.5, "character": { "name": "0men666", "corporation": { "name": "Far-East Corporation", "isNPC": false, "href": "https://api-sisi.testeveonline.com/corporations/98196279/", "id_str": "98196279", "logo": { "32x32": { "href": "https://image.testeveonline.com/Corporation/98196279_32.png" }, "64x64": { "href": "https://image.testeveonline.com/Corporation/98196279_64.png" }, "128x128": { "href": "https://image.testeveonline.com/Corporation/98196279_128.png" }, "256x256": { "href": "https://image.testeveonline.com/Corporation/98196279_256.png" } }, "id": 98196279 }, "isNPC": false, "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "capsuleer": { "href": "https://api-sisi.testeveonline.com/characters/1600000294/capsuleer/" }, "portrait": { "32x32": { "href": "https://image.testeveonline.com/Character/1600000294_32.jpg" }, "64x64": { "href": "https://image.testeveonline.com/Character/1600000294_64.jpg" }, "128x128": { "href": "https://image.testeveonline.com/Character/1600000294_128.jpg" }, "256x256": { "href": "https://image.testeveonline.com/Character/1600000294_256.jpg" } }, "id": 1600000294, "id_str": "1600000294" }, "contact": { "id_str": "1600000294", "href": "https://api-sisi.testeveonline.com/characters/1600000294/", "name": "0men666", "id": 1600000294 }, "href": "https://api-sisi.testeveonline.com/characters/92095466/contacts/1600000294/", "contactType": "Character", "watched": true, "blocked": false } ], "totalCount": 1024, "pageCount_str": "6" }`
	forceFlippableProxyResponse(bodyA, bodyB)
	var accesstoken = "12345"
	var characterid = 67890
	items, err := Character_Contacts(accesstoken, characterid)
	if err != nil {
		t.Errorf("Test_charContacts_paged :: FAIL :: Unexpected error with cause %+v\n", err)
	} else if len(items) != 2 {
		t.Errorf("Test_charContacts_paged :: FAIL :: expected length 2, got '%v'\n", len(items))
	} else {
		t.Logf("Test_charContacts_paged :: Success, response size is %v\n", len(items))
	}
}

/*************************
****Character Fittings****
*************************/

func Test_charFittings(t *testing.T) {
	var body string = `{ "totalCount_str": "1", "items": [ { "description": "THIS IS A TEST FIT DESCRIPTION!", "fittingID_str": "258", "items": [ { "flag": 11, "quantity_str": "1", "type": { "id_str": "518", "href": "http://crest.regner.dev/types/518/", "id": 518, "name": "Basic Gyrostabilizer" }, "flag_str": "11", "quantity": 1 } ], "fittingID": 258, "href": "http://crest.regner.dev/characters/90000001/fittings/258/", "ship": { "id_str": "587", "href": "http://crest.regner.dev/types/587/", "id": 587, "name": "Rifter" }, "name": "THIS IS A TEST FIT NAME!" } ], "pageCount": 1, "pageCount_str": "1", "totalCount": 1 }`
	forceProxyResponse(body)

	var accesstoken string = "12345"
	var characterid int = 67890
	items, err := Character_Fittings(accesstoken, characterid)
	if err != nil {
		t.Errorf("Test_charFittings :: FAIL :: Unexpected error message '%+v'\n", err)
	} else if len(items) != 1 {
		t.Errorf("Test_charFittings :: FAIL :: Expected length 1, got actual length '%v'\n", len(items))
	} else if items[0].FittingID != 258 {
		t.Errorf("Test_charFittings :: FAIL :: Expected fittingID 258, actually got '%v'\n", items[0].FittingID)
	} else {
		t.Logf("Test_charFittings :: Success with length '%+v'\n", len(items))
	}
}

func Test_charFittings_paged(t *testing.T) {
	var bodyA string = `{ "totalCount_str": "2", "items": [ { "description": "THIS IS A TEST FIT DESCRIPTION!", "fittingID_str": "258", "items": [ { "flag": 11, "quantity_str": "1", "type": { "id_str": "518", "href": "http://crest.regner.dev/types/518/", "id": 518, "name": "Basic Gyrostabilizer" }, "flag_str": "11", "quantity": 1 } ], "fittingID": 258, "href": "http://crest.regner.dev/characters/90000001/fittings/258/", "ship": { "id_str": "587", "href": "http://crest.regner.dev/types/587/", "id": 587, "name": "Rifter" }, "name": "THIS IS A TEST FIT NAME!" } ], "pageCount": 2, "pageCount_str": "2", "totalCount": 2, "next" : {"href" : "https://crest-tq.eveonline.com/characters/67890/fittings/?page=2"} }`
	var bodyB string = `{ "totalCount_str": "2", "items": [ { "description": "THIS IS A TEST FIT DESCRIPTION!", "fittingID_str": "258", "items": [ { "flag": 11, "quantity_str": "1", "type": { "id_str": "518", "href": "http://crest.regner.dev/types/518/", "id": 518, "name": "Basic Gyrostabilizer" }, "flag_str": "11", "quantity": 1 } ], "fittingID": 258, "href": "http://crest.regner.dev/characters/90000001/fittings/258/", "ship": { "id_str": "587", "href": "http://crest.regner.dev/types/587/", "id": 587, "name": "Rifter" }, "name": "THIS IS A TEST FIT NAME!" } ], "pageCount": 2, "pageCount_str": "2", "totalCount": 2 }`
	forceFlippableProxyResponse(bodyA, bodyB)

	var accesstoken string = "12345"
	var characterid int = 67890
	items, err := Character_Fittings(accesstoken, characterid)
	if err != nil {
		t.Errorf("Test_charFittings :: FAIL :: Unexpected error message '%+v'\n", err)
	} else if len(items) != 2 {
		t.Errorf("Test_charFittings_paged :: FAIL :: wanted length 2, got '%v'\n", len(items))
	} else {
		t.Logf("Test_charFittings :: Success with length '%+v'\n", len(items))
	}
}

/*************************
****Character Location****
*************************/

func Test_charLocation(t *testing.T) {
	var body = `{ "solarSystem": { "id_str": "30002782", "href": "https://crest-tq.eveonline.com/solarsystems/30002782/", "id": 30002782, "name": "Kamio" } }`
	forceProxyResponse(body)

	var accesstoken string = "12345"
	var characterid int = 67890
	solarSystem, err := Character_Location(accesstoken, characterid)
	if err != nil {
		t.Errorf("Test_charLocation :: FAIL :: got error '%v'\n", err)
	} else if solarSystem.SolarSystem.ID != 30002782 {
		t.Errorf("Test_charLocation :: FAIL :: expected typeid 30002782, got '%v'\n", solarSystem.SolarSystem.ID)
	} else {
		t.Log("Test_charLocation :: Success")
	}
}

/*****************************
***Character Opportunities***
*****************************/

func Test_charOpportunities(t *testing.T) {
	var bodyA = `{ "totalCount_str": "52", "items": [ { "completed": "2016-02-22T19:23:06", "task": { "href": "https://crest-tq.eveonline.com/opportunities/tasks/2/", "id": 2, "id_str": "2" } }, { "completed": "2016-02-21T19:23:06", "task": { "href": "https://crest-tq.eveonline.com/opportunities/tasks/3/", "id": 3, "id_str": "3" } } ], "next": { "href": "https://crest-tq.eveonline.com/characters/6789/characterOpportunitiesRead/" }, "pageCount": 2, "pageCount_str": "2", "totalCount": 4 } `
	var bodyB = `{ "totalCount_str": "52", "items": [ { "completed": "2016-02-19T19:23:06", "task": { "href": "https://crest-tq.eveonline.com/opportunities/tasks/4/", "id": 4, "id_str": "4" } }, { "completed": "2016-02-20T19:23:06", "task": { "href": "https://crest-tq.eveonline.com/opportunities/tasks/5/", "id": 5, "id_str": "5" } } ], "pageCount": 2, "pageCount_str": "2", "totalCount": 4 }`
	forceFlippableProxyResponse(bodyA, bodyB)
	var accesstoken string = "12345"
	var characterid int = 67890
	itemList, err := Character_Opportunities(accesstoken, characterid)
	if err != nil {
		t.Errorf("Test_charOpportunities :: FAIL :: got error %v\n", err)
	} else if len(itemList) != 4 {
		t.Errorf("Test_charOpportunities :: FAIL :: invalid length %v, expected 4\n", len(itemList))
	} else if itemList[0].Completed == "" || itemList[0].Task.ID != 2 {
		t.Errorf("Test_charOpportunities :: FAIL :: bad args for first arg : %+v\n", itemList[0])
	} else {
		t.Log("Test_charOpportunities :: Success")
	}
}
