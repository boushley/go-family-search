package main

// http://play.golang.org/p/7ubbUA1T3R
import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"net/http"
	"os"
	"strings"

	"code.google.com/p/go.net/html"
)

func main() {
	printerConfig := &printer.Config{
		printer.TabIndent & printer.UseSpaces,
		4,
		0,
	}

	resp, err := http.Get("https://familysearch.org/developers/docs/api/fs_json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gatheringUrls []string
	var theUrls [][]string

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
						gatheringUrls = append(gatheringUrls, a.Val)
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
					theUrls = append(theUrls, gatheringUrls)
					gatheringUrls = nil
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
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
		Imports    map[string]bool
	}

	type Enum struct {
		Name   string
		Url    string
		Values []string
	}

	var types []Type
	var enums []Enum
	var doingEnum bool

	processTypeUrl := func(typeUrl string) {
		theType := Type{
			Url:     typeUrl,
			Imports: make(map[string]bool),
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
								if strings.Contains(value, "Link") || strings.Contains(value, "Identifier") {
									value = "map[string]" + strings.TrimSpace(value[8:])
								} else {
									value = "[]" + strings.TrimSpace(value[8:])
								}
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
						t := values[1]
						if t == "boolean" {
							t = "bool"
						} else if t == "float" {
							t = "float32"
						} else if t == "double" {
							t = "float64"
						} else if t == "anyURI" || t == "anyUri" {
							t = "string"
						} else if t == "DateTime" || t == "dateTime" {
							t = "time.Time"
							theType.Imports["time"] = true
						} else if t[:1] != "[" && t != "string" && t != "int" && !strings.HasPrefix(t, "map[string]") {
							t = firstCharacterUpper(t)
						}
						theType.Properties = append(theType.Properties, Property{
							values[0],
							t,
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

	enumUrls := map[string]bool{
		"https://familysearch.org/developers/docs/api/fs/artifactType_json":         true,
		"https://familysearch.org/developers/docs/api/fs/changeObjectModifier_json": true,
		"https://familysearch.org/developers/docs/api/fs/changeObjectType_json":     true,
		"https://familysearch.org/developers/docs/api/fs/changeOperation_json":      true,
		"https://familysearch.org/developers/docs/api/fs/matchStatus_json":          true,
	}
	for _, typeUrl := range theUrls[0] {
		if typeUrl == "https://familysearch.org/developers/docs/api/gx/ResultConfidence_json" {
			continue
		}
		_, doingEnum = enumUrls[typeUrl]
		processTypeUrl(typeUrl)
	}

	doingEnum = true
	for _, typeUrl := range theUrls[1] {
		processTypeUrl(typeUrl)
	}

	fmt.Println("Outputting types")

	for _, t := range types {
		log.Print("Starting on type: " + t.Name)

		typeFile := `
        package models
        `

		if len(t.Imports) > 0 {
			typeFile += `import (
                `

			for key, _ := range t.Imports {
				typeFile += `"` + key + "\"\n"
			}

			typeFile += `
            )
            `
		}

		typeFile += `

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

		fmt.Println("Creating enum: ", t.Name)

		typeFile := `
        package models

        // DataType can be found at: ` + t.Url + `

        type ` + t.Name + ` int

        const (
            _ ` + t.Name + ` = iota
        `

		for _, value := range t.Values {
			typeFile += t.Name + `_` + value[strings.LastIndex(value, `/`)+1:] + ` // ` + value + "\n"
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

		printerConfig.Fprint(outfile, fset, f)
	}

}
