package main

import (
	"fmt"
	// "github.com/gorilla/mux"
	// "github.com/kurt-midas/go-crest/routes"
	// "github.com/kurt-midas/go-crest/sso"
	// "html"
	// "net/http"
	// "net/url"
	// "github.com/kurt-midas/go-crest/calls"
	"github.com/kurt-midas/go-crest/crest"
)

// var app = sso.Conf{"http://localhost:3000/auth/completeHandshake",
// 	"441e626ca5914789ba1e9292e56df4f6",
// "h4M2S4yqvqpwDrw5ncX0krpJNlT51URIjkfZaSiM"}

// ,[]string{"fleetRead", "fleetWrite"}}
// var state = ""

func main() {
	fmt.Println("Hello World!")
	// fmt.Printf("Auth URI: %v\n", app.Auth_uri([]string{"fleetRead", "fleetWrite"}, "123"))

	// allianceList, err := crest.Alliances()
	// fmt.Printf("Alliances: %+v \n\nError: %v\n", allianceList, err)

	// alliance, err := crest.Alliance(498125261)
	// fmt.Printf("Test Alliance : %v\n\nErr : %v\n", alliance, err)

	// dogmaList, err := crest.DogmaAttributes()
	// fmt.Printf("Dogma List : %v\n\nErr : %v\n", dogmaList, err)

	// dogma, err := crest.DogmaAttribute(2111)
	// fmt.Printf("Dogma : %+v\n\nErr : %v\n", dogma, err)

	// list, err := crest.Incursions()
	// fmt.Printf("Result : %v\n\nErr : %v\n", list, err)

	// list, err := crest.InsurancePrices()
	// fmt.Printf("Result : %v\n\nErr : %v\n", list, err)

	// list, err := crest.OpportunitiesGroups()
	// fmt.Printf("Result : %v\n\nErr : %v\n", list, err)

	// item, err := crest.OpportunitiesGroup(112)
	// fmt.Printf("Result : %+v\n\nErr : %+v\n", item, err)

	// list, err := crest.OpportunitiesTasks()
	// fmt.Printf("Result : %+v\n\nErr : %v\n", list, err)

	// item, err := crest.OpportunitiesTask(93)
	// fmt.Printf("Result : %v\n\nErr : %v\n", item, err)

	// res, err := crest.SovereigntyCampaigns()
	// fmt.Printf("Result : %+v\n\nErr : %v\n", res, err)

	// res, err := crest.SovereigntyStructures()
	// fmt.Printf("Result : %+v\n\nErr : %v\n", res, err)

	// res, pagecount, err := crest.InventoryTypes_Page(2)
	// fmt.Printf("Result : %v\n\nPages: %v\n\nErr : %v\n", res, pagecount, err)

	res, err := crest.InventoryType(12005)
	fmt.Printf("Result : %+v\n\nErr : %v\n", res, err)

	// routes.HelloWorldRoute()
	// calls.AvoidWarning()
	// crest.AvoidWarning()

	// fmt.Println(app.GetSSORedirectURL())

	// r := mux.NewRouter()
	// r.HandleFunc("/", getSSOUrl)
	// r.HandleFunc("/auth/completeHandshake", catchHandshake)
	// http.Handle("/", r)
	// http.ListenAndServe(":3000", nil)
}

/*
func getSSOUrl(w http.ResponseWriter, r *http.Request) {
	scopes := []string{"123", "456"}
	ssoUrl, urlErr := app.Auth_uri(scopes, "789")
	if urlErr != nil {
		fmt.Fprintf(w, urlErr.Error())
		return
	}
	fmt.Println("ssoUrl", ssoUrl.String())
	fmt.Println("state", state)
	fmt.Fprint(w, ssoUrl.String())
}

func catchHandshake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside 'catchHandshake'")
	code, _ := app.HandleOauthResponse(r)
	response, handshakeErr := app.AccessToken(code)
	if handshakeErr != nil {
		fmt.Printf("Got error while catching handshake :: %v\n", handshakeErr.Error())
		fmt.Fprintf(w, handshakeErr.Error())
		return
	}
	fmt.Printf("Successful response from app.CompleteHandshake :: %+v\n", response)
	fmt.Fprintf(w, "Successful response from app.CompleteHandshake")
	return
}*/

/*func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, World!")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}*/
