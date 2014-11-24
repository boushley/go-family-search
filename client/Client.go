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
	root   Endpoint
	http   *http.Client
	config Config
	tree   *models.FamilySearch
}

func NewClient(root Endpoint, conf Config) *Client {
	c := Client{}
	c.root = root
	c.http = &http.Client{}
	c.config = conf
	return &c
}

func (c *Client) Initialize() error {
	c.getRootTree()
	return nil
}

func (c *Client) getRootTree() {
	res := c.makeRequest("GET", c.root.TreeRoot, nil)

	t := &models.FamilySearch{}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	handleError(err)

	err = json.Unmarshal(bodyBytes, t)
	handleError(err)

	if t == nil {
		log.Fatal("Unable to parse root config tree.")
	}

	c.tree = t
}

func (c *Client) makeRequest(method, url string, body io.Reader) (res *http.Response) {
	req, err := http.NewRequest(method, url, body)
	handleError(err)

	req.Header.Set("Accept", "application/x-fs-v1+json")
	req.Header.Set("Content-Type", "application/json")

	res, err = c.http.Do(req)
	handleError(err)

	return
}

func (c *Client) makeAuthenticatedRequest(method, url, accessToken string, body io.Reader) (res *http.Response) {
	req, err := http.NewRequest(method, url, body)
	handleError(err)

	req.Header.Set("Accept", "application/x-fs-v1+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	res, err = c.http.Do(req)
	handleError(err)
	return
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
