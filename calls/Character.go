package calls

import (
	"fmt"
	"github.com/kurt-midas/go-crest/utils"
	"net/url"
)

/* TODO
Contacts
	Character_GetContacts
	Character_WriteContact
	Contact error handling and stuff
Contact
	PUT
	DELETE
Fittings
	GET
	POST
Fitting
	GET
	DELETE
Opportunities
	GET
Waypoints
	POST
Location
	GET
*/

func Call_CharContacts(accesstoken string) (error, []CharContactItems) {
	// utils.CrestDomain + "path?"
	reqErr, req := utils.BuildCrestCallRequest("GET", "rawurl", accesstoken, url.Values{})
	// var returnObj = new(CharContactsFormat)
	var returnObj = new(CharContactsFormat)
	if reqErr != nil {
		return reqErr, returnObj.Items
	}
	callErr := utils.RemotePageableCall(req, returnObj)
	if callErr != nil {
		return callErr, returnObj.Items
	}
	return callErr, returnObj.Items
}

type CharContactsFormat struct {
	PageCount      int                `json:"pageCount"`
	PageCount_str  string             `json:"pageCount_str"`
	TotalCount     int                `json:"totalCount"`
	TotalCount_str string             `json:"totalCount_str"`
	Items          []CharContactItems `json:"items"`
	Next           utils.SubHref      `json:"next"`
}

func (c CharContactsFormat) GetCurrentPageNumber() int { return c.PageCount }
func (c CharContactsFormat) GetTotalPageCount() int    { return c.TotalCount }
func (c CharContactsFormat) GetNext() utils.SubHref    { return c.Next }
func (c CharContactsFormat) GetItems() utils.Any       { return c.Items }
func (c CharContactsFormat) SetItems(a utils.Any) {
	if val, ok := a.([]CharContactItems); ok {
		c.Items = val
	} else {
		fmt.Println("CharContactsFormat setItems failure, wrong type")
	}
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
	Capsuleer utils.SubHref   `json:"capsuleer"`
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

/**
 * Should go to a SharedStructs.go file, maybe even its own package
 */

type ImgStruct struct {
	Size32  utils.SubHref `json:"32x32"`
	Size64  utils.SubHref `json:"64x64"`
	Size128 utils.SubHref `json:"128x128"`
	Size256 utils.SubHref `json:"256x256"`
}
