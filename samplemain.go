package main

import "github.com/kurt-midas/go-crest/sso"
import "fmt"
import "net/http"

var app = sso.Conf{"http://localhost:3000/auth/completeHandshake", //callback URI
	"441e626ca5914789ba1e9292e56df4f6",         //client id
	"h4M2S4yqvqpwDrw5ncX0krpJNlT51URIjkfZaSiM"} //client secret

/*func main() {
	http.HandleFunc("/", getSSOUrl)
	http.HandleFunc("/auth/completeHandshake", catchHandshake)
	fmt.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}*/

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
