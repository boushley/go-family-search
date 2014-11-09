package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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

func (c *Client) GetAuthorizationEndpoint(landingUrl string) string {
	baseUrl := c.tree.Collections[0].Links["http://oauth.net/core/2.0/endpoint/authorize"].Href
	values := url.Values{}
	values.Add("response_type", "code")
	values.Add("redirect_uri", landingUrl)
	values.Add("client_id", c.config.ClientId)
	return baseUrl + "?" + values.Encode()
}

func (c *Client) GetTokenEndpoint() string {
	return c.tree.Collections[0].Links["http://oauth.net/core/2.0/endpoint/token"].Href
}

func (c *Client) GetTokenFromCode(code string) (token string) {
	tokenUrl, err := url.Parse(c.GetTokenEndpoint())
	handleError(err)
	query := tokenUrl.Query()
	query.Add("grant_type", "authorization_code")
	query.Add("code", code)
	query.Add("client_id", c.config.ClientId)
	tokenUrl.RawQuery = query.Encode()

	res := c.makeRequest("POST", tokenUrl.String(), nil)
	bodyBytes, err := ioutil.ReadAll(res.Body)
	handleError(err)
	fmt.Println("Response Received: " + string(bodyBytes))

	oauth := &models.OAuthTokenResponse{}
	err = json.Unmarshal(bodyBytes, oauth)
	handleError(err)

	if oauth.Error != "" {
		fmt.Println("Error getting access token: " + oauth.Error + " Description: " + oauth.ErrorDescription)
	}

	return oauth.AccessToken
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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
