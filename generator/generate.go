package main

// http://play.golang.org/p/7ubbUA1T3R
import (
	"code.google.com/p/go.net/html"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"net/http"
	"os"
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

	firstCharacterUpper := func(name string) string {
		return strings.ToUpper(name[:1]) + name[1:]
	}

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

	type Type struct {
		Name       string
		Url        string
		Properties []Property
	}

	type Enum struct {
		Name   string
		Url    string
		Values []string
	}

	var types []Type
	var enums []Enum
	var doingEnum bool

	enumUrls := map[string]int{
		"https://familysearch.org/developers/docs/api/fs/artifactType_json":         1,
		"https://familysearch.org/developers/docs/api/fs/changeObjectModifier_json": 1,
		"https://familysearch.org/developers/docs/api/fs/changeObjectType_json":     1,
		"https://familysearch.org/developers/docs/api/fs/changeOperation_json":      1,
		"https://familysearch.org/developers/docs/api/fs/matchStatus_json":          1,
		"https://familysearch.org/developers/docs/api/gx/ResultConfidence_json":     1,
	}
	for _, typeUrl := range typeUrls {
		_, doingEnum = enumUrls[typeUrl]

		theType := Type{
			Url: typeUrl,
		}
		theEnum := Enum{
			Url: typeUrl,
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

					if doingEnum {
						theEnum.Values = append(theEnum.Values, strings.TrimSpace(values[0]))
					} else if len(values) == 4 {
						if strings.TrimSpace(values[0]) == "" || strings.TrimSpace(values[1]) == "" {
							log.Fatal("Missing key value: ", values)
						}
						theType.Properties = append(theType.Properties, Property{
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
			if n.Type == html.ElementNode && n.Data == "h1" {
				title := gatherText(n)
				if strings.TrimSpace(title) != "" {
                    name := strings.TrimSpace(title[:strings.Index(title, ` `)])
					theType.Name = name
                    theEnum.Name = firstCharacterUpper(name)
				}
			} else if n.Type == html.ElementNode && n.Data == "table" {
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

		if doingEnum {
			enums = append(enums, theEnum)
		} else {
			types = append(types, theType)
		}
	}

	fmt.Println("Outputting types")

	for _, t := range types {

		typeFile := `
        package models

        // DataType can be found at: ` + t.Url + `

        type ` + t.Name + ` struct {
        `

		for _, p := range t.Properties {
			propType := p.Type
			if propType != "string" {
				propType = firstCharacterUpper(propType)
			}

			typeFile += firstCharacterUpper(p.Name) + ` ` + p.Type + " `json:\"" + p.Name + "\"`\n"
		}
		typeFile += ` } `

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", typeFile, 0)
		if err != nil {
			log.Fatal(err)
		}

		outfile, err := os.Create("../models/" + t.Name + ".go")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Writing out file: ", outfile.Name())

		printer.Fprint(outfile, fset, f)
	}

	fmt.Println("Outputting enums")

	for _, t := range enums {

		typeFile := `
        package models

        // DataType can be found at: ` + t.Url + `

        type ` + t.Name + ` int

        const (
            _ ` + t.Name + ` = iota
        `

		for _, value := range t.Values {
            typeFile += value[strings.LastIndex(value, `/`):] + ` // ` + value + "\n"
		}
		typeFile += ` ) `

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "", typeFile, 0)
		if err != nil {
			log.Fatal(err)
		}

		outfile, err := os.Create("../models/" + t.Name + ".go")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Writing out file: ", outfile.Name())

		printer.Fprint(outfile, fset, f)
	}

}
