package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/boushley/go-family-search/models"
)

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

	oauth := &models.OAuthTokenResponse{}
	err = json.Unmarshal(bodyBytes, oauth)
	handleError(err)

	if oauth.Error != "" {
		fmt.Println("Error getting access token: " + oauth.Error + " Description: " + oauth.ErrorDescription)
	}

	return oauth.AccessToken
}
