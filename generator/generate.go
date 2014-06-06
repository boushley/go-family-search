package main

// http://play.golang.org/p/7ubbUA1T3R
import (
	//"go/parser"
	//"go/printer"
	//"go/token"
	"code.google.com/p/go.net/html"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://familysearch.org/developers/docs/api/fs_json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var typeUrls []string

	var gatherText func(n *html.Node) string
	gatherText = func(n *html.Node) (result string) {

		if n.Type == html.TextNode {
			return n.Data
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result += gatherText(c)
		}

		return
	}

	var findTypes func(n *html.Node)
	findTypes = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.Data == "div" {
				link := c.FirstChild
				for _, a := range link.Attr {
					if a.Key == "href" {
						typeUrls = append(typeUrls, a.Val)
					}
				}
			}
		}
	}
	var findTypesContainer func(n *html.Node)
	findTypesContainer = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "datatypes" {
					findTypes(n)
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if len(typeUrls) > 0 {
				return
			}
			findTypesContainer(c)
		}
	}
	findTypesContainer(doc)

	type Property struct {
		Name        string
		Type        string
		Description string
	}

	skipUrls := map[string]int{
		"https://familysearch.org/developers/docs/api/fs/artifactType_json":         1,
		"https://familysearch.org/developers/docs/api/fs/changeObjectModifier_json": 1,
		"https://familysearch.org/developers/docs/api/fs/changeObjectType_json":     1,
		"https://familysearch.org/developers/docs/api/fs/changeOperation_json":      1,
		"https://familysearch.org/developers/docs/api/fs/matchStatus_json":          1,
		"https://familysearch.org/developers/docs/api/gx/ResultConfidence_json":     1,
	}
	for _, typeUrl := range typeUrls {
		if _, ok := skipUrls[typeUrl]; ok {
			fmt.Println("Skipping format info for: ", typeUrl)
			continue
		}

		fmt.Println("Parsing format info for: ", typeUrl)
		resp, err := http.Get(typeUrl)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var props []Property
		var values []string

		var findProperties func(n *html.Node)
		findProperties = func(n *html.Node) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "tr" {
					values = nil

					for cell := c.FirstChild; cell != nil; cell = cell.NextSibling {
						if cell.Type == html.ElementNode && cell.Data == "td" {
							value := strings.TrimSpace(gatherText(cell))
							if strings.HasPrefix(value, "array of") {
								value = "[]" + strings.TrimSpace(value[8:])
							}
							values = append(values, value)
						}
					}

					if len(values) == 4 {
						if strings.TrimSpace(values[0]) == "" || strings.TrimSpace(values[1]) == "" {
							log.Fatal("Missing key value: ", values)
						}
						props = append(props, Property{
							values[0],
							values[1],
							values[2],
						})
					} else if len(values) != 0 {
						log.Fatal("Skipping values: ", values)
					}
				}
			}
		}
		var findPropertiesTable func(n *html.Node)
		findPropertiesTable = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "table" {
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "table table-striped" {
						for c := n.FirstChild; c != nil; c = c.NextSibling {
							if c.Type == html.ElementNode && c.Data == "tbody" {
								findProperties(c)
							}
						}
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				findPropertiesTable(c)
			}
		}
		findPropertiesTable(doc)

		fmt.Println(props)
	}

}
