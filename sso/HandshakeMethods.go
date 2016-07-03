package sso

import (
	"bytes"
	"encoding/base64"
	// "encoding/json"
	"fmt"
	"github.com/kurt-midas/go-crest/utils"
	// "io/ioutil"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

/*
TODO:
	Error handling : successful but invalid responses from CREST
*/

//trying to do a bunch of fixed scopes, isn't really working
const (
	FleetRead  string = "fleetRead"
	FleetWrite string = "fleetWrite"
)

/**********
Response type for token stuff
**********/
type OauthTokenResponse struct {
	AuthToken    string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

/**********
Main type, required for everything else
**********/
type Conf struct {
	CallbackURI  string
	ClientID     string
	ClientSecret string
	// Scopes       []string
}

/**
 *
 * @param  (c Conf)
 * @return (ssoUrl string, state string)   [description]
 */
func (c Conf) Auth_uri(scopes []string, state string) (url.URL, error) {
	// fmt.Println("Auth_uri :: conf is", c)

	ssoUrl, err := url.Parse("https://login.eveonline.com/oauth/authorize")
	if err != nil {
		return *ssoUrl, errors.New("Fatal error parsing SSO url")
	}
	parameters := url.Values{}
	parameters.Add("response_type", "code")
	parameters.Add("redirect_uri", c.CallbackURI)
	parameters.Add("client_id", c.ClientID)
	parameters.Add("scope", strings.Join(scopes, " "))
	parameters.Add("state", state)
	ssoUrl.RawQuery = parameters.Encode()

	return *ssoUrl, nil
}

//Parses callback from Oauth response into "code" and "state"
//Utility method, no real functionality. Probably shouldn't exist
func (c Conf) HandleOauthResponse(r *http.Request) (string, string) {
	// fmt.Println("PublicMethods :: CompleteHandshake :: Query is : ", r.URL.Query())
	return r.URL.Query().Get("code"), r.URL.Query().Get("state")
	// return c.exchangeToken(r.URL.Query().Get("code"), "auth")
}

func (c Conf) AccessToken(authtoken string) (OauthTokenResponse, error) {
	return c.exchangeToken(authtoken, "auth")
}

func (c Conf) RefreshToken(refreshtoken string) (OauthTokenResponse, error) {
	return c.exchangeToken(refreshtoken, "refresh")
}

//valid token types are "auth" and "refresh"
func (c Conf) exchangeToken(token string, tokenType string) (OauthTokenResponse, error) {
	var method string = "POST"
	var urlStr string = "https://login-tq.eveonline.com/oauth/token"
	data := url.Values{}
	if tokenType == "auth" {
		data.Add("grant_type", "authorization_code")
		data.Add("code", token)
	} else if tokenType == "refresh" {
		data.Add("grant_type", "refresh_token")
		data.Add("refresh_token", token)
	} else {
		badUseErr := errors.New("Error Exchanging Token :: Only supported types are 'auth' and 'refresh', instead got " + tokenType)
		fmt.Printf("PublicMethods :: ExchangeToken :: %v\n", badUseErr.Error())
		return OauthTokenResponse{}, badUseErr
	}
	req, reqErr := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	// checkErr(reqErr)
	if reqErr != nil {
		return OauthTokenResponse{}, reqErr
	}
	var concat = c.ClientID + ":" + c.ClientSecret
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(concat)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var respStruct = new(OauthTokenResponse)
	callErr := utils.RemoteCall(req, respStruct)
	if callErr != nil {
		return *respStruct, callErr
	}
	return *respStruct, nil
}
