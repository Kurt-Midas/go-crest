#GO-CREST
This package contains an importable golang library which can handle CREST authentication and calls. 

The CREST API, aka the Carbon RESTful API, is a service maintained by CCP (the company behind the MMO Eve Online) to aid third-party developers in developing apps using in-game data. The purpose of this library is to abstract some of the access details to help developers do more, better and faster. 

### Quick use notes ###
Please note that use of this package implies acceptance of CCP's Third-Party Developer agreement. This agreement bans most monetization and all scams. 

You may use, modify, and share this code however you want.

If you enjoy the library, the best "thank you" is to submit suggestions or requests. 

### Design Disclaimer ###
This library uses very different handling than other CREST libraries like pycrest or fuzzy's PHP library. Specifically, this library does NOT walk the api and dynamically generate endpoints or response structs. While this makes the library fragile to changes in CREST, the rough equivalent of a walkable dynamically generated API would be a library where all responses are of type map[string]T. This discards the primary strength of a typed language: ease of coding. Static structs where all fields are of known type are immeasurably easier to work with. 

If you notice that behavior no longer appears to work as intended, please submit a ticket. It should be possible for you to modify the URL and structs yourself and "go install" the changes if I don't respond quickly. 

### Use Disclaimer ###
I created this library primarily to centralize my own CREST access. I do not recommend using this package if you are not me. If you still want to use this package, please reach out to me (github is probably the most consistent way) and consider the following:

This library is incomplete. Several important calls have been added but I have not attempted to add or maintain everything. If you notice a certain call is NOT supported, you may add it yourself using existing constructs or contact me/submit a ticket. 

Design may change at any time. Code may break without warning. I haven't finalized how I want to handle packages or configuration yet. I'll try to provide warning if I know you exist. 

There may be errors. I've verified every call works under optimal circumstances but certain problems haven't been tested or addressed (see https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779#.1sdu4ns4w ). Tell me if there's a problem so I can bump it on my priority list. There's also pretty bad debug logging, another low-priority thing. 

Types have marginally inconsistent behavior around semi-duplicate fields. For example, CREST sends "id" int and "id_str" string in the same response. Some of my types have both, some only have int. I don't have a problem and can't know you have one unless you contact me. 

# Tutorial #
## How to use ##
I haven't tested "go get" yet. I'll change this once I have, but you can download the package into your gopath and install it without too much trouble. 
## Calls ##
```go
package main

import "github.com/kurt-midas/go-crest/crest"
import "fmt"

func main(){
	res, err := crest.InventoryType(12005)
	fmt.Printf("Result : %+v\n\nErr : %v\n", res, err)
}
```
All API methods and response types are attached directly to the "crest" package. 

##SSO Setup
Certain calls require an "access token." Generating this requires use of the SSO. A simple program can be found below:
```
package main

import "github.com/kurt-midas/go-crest/sso"
import "github.com/kurt-midas/go-crest/crest"
import "fmt"
import "net/http"

var app = sso.Conf{"http://localhost:3000/auth/completeHandshake", //callback URI
	"441e626ca5914789ba1e9292e56df4f6", //client id
	"h4M2S4yqvqpwDrw5ncX0krpJNlT51URIjkfZaSiM"} //client secret

func main() {
	http.HandleFunc("/", getSSOUrl)
	http.HandleFunc("/auth/completeHandshake", catchHandshake)
	fmt.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func getSSOUrl(w http.ResponseWriter, r *http.Request) {
	scopes := []string{"fleetRead", "fleetWrite"}
	ssoUrl, urlErr := app.Auth_uri(scopes, "789") //create, store, and handle a unique state (time factor etc)
	if urlErr != nil {
		fmt.Fprintf(w, urlErr.Error())
		return
	}
	fmt.Println("ssoUrl", ssoUrl.String())
	fmt.Fprint(w, ssoUrl.String())
}

func catchHandshake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside 'catchHandshake'")
	code, _ := app.HandleOauthResponse(r) //ignoring state, which is bad
	response, handshakeErr := app.AccessToken(code)
	if handshakeErr != nil {
		fmt.Printf("Got error while catching handshake :: %v\n", handshakeErr.Error())
		fmt.Fprintf(w, handshakeErr.Error())
		return
	}
	fmt.Printf("Successful response from app.CompleteHandshake :: %+v\n", response)
	fmt.Fprintf(w, "Successful response from app.CompleteHandshake")
}
```
This is oversimplified, but it shows the general flow. A user lands on localhost:3000/ and is served (but not redirected) to an SSO url. On logging in, that user is redirected back to http://localhost:3000/auth/completeHandshake where the app exchanges the authorization token for an access token. Please note that (currently) the above code represents a ccp-registered application and can be copy-pasted. 

Please register your own app. These details should technically be private. 

##What the SSO actually is
Deeper information about the SSO and what exactly this is doing behind the scenes can be found on the kinda-official documentation.

http://eveonline-third-party-documentation.readthedocs.io/en/latest/index.html

To summarize correctly implemented SSO flow:

1.	Your app redirects a user to your generated SSO url (controlled by CCP)

2.	When the user logs in, CCP's SSO site will redirect the user back to your registered callback URL. There will be two query parameters attached to the request: "code" and "state"

3.	The "code" parameter is an authorization code. Your app should exchange this authorization code for an access and refresh token.

4.	The access token is used to access authed endpoints. Access tokens currently last for 1200 seconds (20 minutes) before they become invalid

5.	The refresh token is used to generate another access token. Refresh tokens do not expire. 


To summarize important points about the SSO:

*	You MUST register your app with CCP. You MUST have a working callback URL (localhost:PORT is valid). 

*	Refresh tokens do not expire. Users can only invalidate refresh tokens through CCP's support site. Most do not know this. Be very, VERY careful about how they are handled.

*	Violating the CCP third-party developer agreement is a bannable offense. You have been warned. 
