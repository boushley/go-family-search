package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boushley/go-family-search/models"
)

type Client struct {
	root Endpoint
	http *http.Client

	tokenEndpoint     string
	authorizeEndpoint string
}

func NewClient(root Endpoint) Client {
	c := Client{}
	c.root = root
	c.http = &http.Client{}
	return c
}

func (c Client) Initialize() error {
	tree := c.getRootTree()

	c.tokenEndpoint = tree.Collections[0].Links["http://oauth.net/core/2.0/endpoint/token"].Href
	c.authorizeEndpoint = tree.Collections[0].Links["http://oauth.net/core/2.0/endpoint/authorize"].Href

	log.Println(c.tokenEndpoint)
	log.Println(c.authorizeEndpoint)

	return nil
}

func (c Client) getRootTree() (tree *models.FamilySearch) {
	res := c.makeRequest("GET", c.root.TreeRoot, nil)

	tree = &models.FamilySearch{}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	handleError(err)

	err = json.Unmarshal(bodyBytes, tree)
	handleError(err)

	return
}

func (c Client) makeRequest(method, url string, body io.Reader) (res *http.Response) {
	req, err := http.NewRequest(method, url, body)
	handleError(err)

	req.Header.Set("Accept", "application/x-fs-v1+json")
	req.Header.Set("Content-Type", "application/json")

	res, err = c.http.Do(req)
	handleError(err)

	return
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
