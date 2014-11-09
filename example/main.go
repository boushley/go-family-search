package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)
import "github.com/boushley/go-family-search/client"

var cli *client.Client
var config Configuration

func main() {
	loadConfig()
	initClient()

	http.HandleFunc("/login", Login)
	http.HandleFunc("/familySearch/auth", OAuthReturn)
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

	fmt.Println("Got Code: " + code + " token: " + token)
	// Set token in a session store or a cookie
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() {
	file, err := os.Open("private-conf.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func initClient() {
	conf := client.Config{config.DeveloperKey}
	cli = client.NewClient(client.Sandbox, conf)
	cli.Initialize()
}
