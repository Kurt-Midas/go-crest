package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kurt-midas/go-crest/routes"
	"github.com/kurt-midas/go-crest/sso"
	// "html"
	"net/http"
	// "net/url"
	"github.com/kurt-midas/go-crest/calls"
	"github.com/kurt-midas/go-crest/crest"
)

var app = sso.Conf{"http://localhost:3000/auth/completeHandshake",
	"441e626ca5914789ba1e9292e56df4f6",
	"h4M2S4yqvqpwDrw5ncX0krpJNlT51URIjkfZaSiM",
	[]string{"fleetRead", "fleetWrite"}}
var state = ""

func main() {
	fmt.Println("Hello World!")
	routes.HelloWorldRoute()
	calls.AvoidWarning()
	crest.AvoidWarning()

	// fmt.Println(app.GetSSORedirectURL())

	r := mux.NewRouter()
	r.HandleFunc("/", getSSOUrl)
	r.HandleFunc("/auth/completeHandshake", catchHandshake)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func getSSOUrl(w http.ResponseWriter, r *http.Request) {
	urlErr, ssoUrl, state := app.GetSSORedirectURL()
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
	handshakeErr, response := app.CompleteHandshake(w, r)
	if handshakeErr != nil {
		fmt.Printf("Got error while catching handshake :: %v\n", handshakeErr.Error())
		fmt.Fprintf(w, handshakeErr.Error())
		return
	}
	fmt.Printf("Successful response from app.CompleteHandshake :: %+v\n", response)
	fmt.Fprintf(w, "Successful response from app.CompleteHandshake")
	return
}

/*func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, World!")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}*/
