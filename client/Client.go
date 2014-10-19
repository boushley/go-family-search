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
}

func NewClient(root Endpoint) Client {
	httpClient := &http.Client{}
	c := Client{root, httpClient}
	return c
}

func (c Client) Initialize() error {
	tree := c.getRootTree()
	treeBytes, _ := json.Marshal(tree)
	log.Print(string(treeBytes))

	return nil
}

func (c Client) getRootTree() (tree *models.FamilySearch) {
	res := c.makeRequest("GET", c.root.TreeRoot, nil)

	tree = &models.FamilySearch{}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	log.Print(string(bodyBytes))
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
