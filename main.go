package main

import (
	"github.com/boushley/go-family-search/client"
)

func main() {
	c := client.NewClient(client.Sandbox)
	c.Initialize()
}
