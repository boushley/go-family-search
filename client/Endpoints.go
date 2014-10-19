package client

type Endpoint struct {
	TreeRoot string
}

var (
	Prod    Endpoint = Endpoint{"https://familysearch.org/platform/collections/tree"}
	Beta    Endpoint = Endpoint{"https://beta.familysearch.org/platform/collections/tree"}
	Sandbox Endpoint = Endpoint{"https://sandbox.familysearch.org/platform/collections/tree"}
)
