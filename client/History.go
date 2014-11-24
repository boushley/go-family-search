package client

import (
	"encoding/json"
	"fmt"
	"github.com/boushley/go-family-search/models"
	"io/ioutil"
	"log"
)

func (c *Client) GetCurrentUserHistory(accessToken string) []models.SourceDescription {
	url := c.getCurrentUserHistoryPath()
	res := c.makeAuthenticatedRequest("GET", url, accessToken, nil)

	t := &models.FamilySearch{}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))
	handleError(err)

	err = json.Unmarshal(bodyBytes, t)
	handleError(err)

	if t == nil {
		log.Fatal("Unable to parse response from current user history")
	}

	history := t.SourceDescriptions

	return history
}

func (c *Client) getCurrentUserHistoryPath() string {
	return c.tree.Collections[0].Links["current-user-history"].Href
}
