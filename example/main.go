package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)
import "github.com/boushley/go-family-search/client"
import "github.com/gorilla/sessions"

var cli *client.Client
var config Configuration
var store *sessions.CookieStore

func main() {
	loadConfig()
	initClient()

	store = sessions.NewCookieStore([]byte("aoiajdsio@309FA)023-aglkj"))

	http.HandleFunc("/login", Login)
	http.HandleFunc("/familySearch/auth", OAuthReturn)
	http.HandleFunc("/recent-people", ShowRecentPeople)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func Login(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Login")
	authEndpoint := cli.GetAuthorizationEndpoint("http://localhost:8080/familySearch/auth")
	http.Redirect(rw, r, authEndpoint, http.StatusTemporaryRedirect)
}

func OAuthReturn(rw http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	token := cli.GetTokenFromCode(code)

	session, err := store.Get(r, "auth")
	handleError(err)

	session.Values["accessToken"] = token
	err = session.Save(r, rw)
	handleError(err)

	http.Redirect(rw, r, "/recent-people", http.StatusTemporaryRedirect)
}

func ShowRecentPeople(w http.ResponseWriter, r *http.Request) {
	history := cli.GetCurrentUserHistory(getToken(r))
	fmt.Println(history)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() {
	file, err := os.Open("private-conf.json")
	handleError(err)

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	handleError(err)
}

func initClient() {
	conf := client.Config{config.DeveloperKey}
	cli = client.NewClient(client.Sandbox, conf)
	cli.Initialize()
}

func getToken(r *http.Request) string {
	session, err := store.Get(r, "auth")
	handleError(err)
	return session.Values["accessToken"].(string)
}
