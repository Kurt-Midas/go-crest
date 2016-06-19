package sso

import (
	"bytes"
	"encoding/base64"
	// "encoding/json"
	"fmt"
	"github.com/kurt-midas/go-crest/utils"
	// "io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
TODO:
	Make State a time factor
	Do something with State, it needs to be verified somehow
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
Error for invalid usage. Ex: someone tries using an unexpected scope
**********/
type InvalidUsageError struct {
	Message string
	Cause   string
}

func (t InvalidUsageError) Error() string {
	return fmt.Sprintf("SSO Error :: %v :: %v", t.Message, t.Cause)
}

/**********
Main type, required for everything else
**********/
type Conf struct {
	CallbackURI  string
	ClientID     string
	ClientSecret string
	Scopes       []string
}

/**
 * Returns the SSO url to send the user to and the state variable
 * @param  (c Conf)
 * @return (ssoUrl string, state string)   [description]
 */
func (c Conf) GetSSORedirectURL() (error, url.URL, string) {
	fmt.Println("GetSSORedirectURL :: conf is", c)
	var state string = "1234567890" //TODO: add time factor

	ssoUrl, err := url.Parse("https://login.eveonline.com/oauth/authorize")
	if err != nil {
		return nil, *ssoUrl, ""
	}
	// checkErr(err)
	parameters := url.Values{}
	parameters.Add("response_type", "code")
	parameters.Add("redirect_uri", c.CallbackURI)
	parameters.Add("client_id", c.ClientID)
	parameters.Add("scope", strings.Join(c.Scopes, " "))
	parameters.Add("state", state)
	ssoUrl.RawQuery = parameters.Encode()

	// fmt.Println("Query Encode :: " + ssoUrl.Query().Encode())
	// fmt.Println("ssoUrl :: " + ssoUrl.String())
	// fmt.Println("Path :: " + ssoUrl.Path)
	// fmt.Println("Raw Path :: " + ssoUrl.RawPath)
	// fmt.Println("Request URI :: " + ssoUrl.RequestURI())
	// fmt.Println("Escaped Path :: " + ssoUrl.EscapedPath())
	return nil, *ssoUrl, state
}

func (c Conf) CompleteHandshake(w http.ResponseWriter, r *http.Request) (error, OauthTokenResponse) {
	//code and state
	// fmt.Println("PublicMethods :: CompleteHandshake :: Query is : ", r.URL.Query())
	return c.exchangeToken(r.URL.Query().Get("code"), "auth")
}

func (c Conf) ExchangeAuthToken(token string) (error, OauthTokenResponse) {
	return c.exchangeToken(token, "auth")
}

func (c Conf) RefreshToken(token string) (error, OauthTokenResponse) {
	return c.exchangeToken(token, "refresh")
}

//valid token types are "auth" and "refresh"
func (c Conf) exchangeToken(token string, tokenType string) (error, OauthTokenResponse) {
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
		badUseErr := InvalidUsageError{"Error Exchanging Token", "Only supported types are 'auth' and 'refres', instead got " + tokenType}
		fmt.Printf("PublicMethods :: ExchangeToken :: %v\n", badUseErr.Error())
		return badUseErr, OauthTokenResponse{}
	}
	req, reqErr := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
	// checkErr(reqErr)
	if reqErr != nil {
		return reqErr, OauthTokenResponse{}
	}
	var concat = c.ClientID + ":" + c.ClientSecret
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(concat)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var respStruct = new(OauthTokenResponse)
	callErr := utils.RemoteCall(req, respStruct)
	if callErr != nil {
		return callErr, *respStruct
	}
	return nil, *respStruct
}
