package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

//Reference: https://skarlso.github.io/2016/06/12/google-signin-with-go/

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

// User is a retrieved and authentiacted user.
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender        string `json:"gender"`
}

var cred = Credentials{"530207568765-e3co60hkd7r8fs36bm2acghdqtckbrh5.apps.googleusercontent.com", "bGDv7imH0hyXq1QA7RXuRBnQ"}
var conf *oauth2.Config
var state string

// var store = sessions.NewCookieStore([]byte("secret"))

func init() {
}

func main() {
	initHandler()

}

func initHandler() {
	route := mux.NewRouter()
	route.HandleFunc("/login", loginHandler).Methods("GET")
	route.HandleFunc("/oauth2callback", callbackFunction).Methods("POST")
	// route.HandleFunc("/reports/{id}", h.getReport).Methods("GET")

	serverHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}(route)
	http.ListenAndServe(":9090", serverHandler)
}

func callbackFunction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	codeVal := vars["code"]

	fmt.Printf("%s", codeVal)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	endpoint := oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://www.googleapis.com/oauth2/v3/token",
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:9090/oauth2callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: endpoint,
	}

	fmt.Printf("config completed")
	state := "testingstate"
	fmt.Printf(conf.AuthCodeURL(state))
}
