package controllers

import (
	"log"

	"github.com/revel/revel"
)
import "github.com/boushley/go-family-search/client"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Login() revel.Result {
	privateConfig, err := revel.LoadConfig("private.conf")
	if err != nil {
		log.Fatal(err)
	}

	key, found := privateConfig.String("app.private.developerKey")
	if !found {
		log.Fatal("Missing app.private.developerKey config, did you copy conf/private.conf.example and add your key")
	}

	conf := client.Config{key}
	cli := client.NewClient(client.Sandbox, conf)
	cli.Initialize()

	authEndpoint := cli.GetAuthorizationEndpoint("http://localhost:8080/familySearch/auth")

	return c.Redirect(authEndpoint)
}

func (c App) Authorize() revel.Result {
	privateConfig, err := revel.LoadConfig("private.conf")
	if err != nil {
		log.Fatal(err)
	}

	key, found := privateConfig.String("app.private.developerKey")
	if !found {
		log.Fatal("Missing app.private.developerKey config, did you copy conf/private.conf.example and add your key")
	}

	conf := client.Config{key}
	cli := client.NewClient(client.Sandbox, conf)
	cli.Initialize()

	return c.Render()
}
