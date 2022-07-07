package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const (
	pkgDocLink     = "https://pkg.go.dev/k8s.io/component-helpers/storage/volume"
	sectionToParse = `<section class="Documentation-constants">`
)

func getPackageDocContent() ([]byte, error) {
	resp, err := http.Get(pkgDocLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// TODO: Handle non OK response codes
	return ioutil.ReadAll(resp.Body)
}

func parseConstSection(doc string) []string {
	tkn := html.NewTokenizer(strings.NewReader(doc))
	var vals []string
	var isPkgConstSection bool
	for {

		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return vals

		case tt == html.StartTagToken:
			if isPkgConstSection {
				continue
			}
			t := tkn.Token()
			isPkgConstSection = t.String() == sectionToParse

		case tt == html.EndTagToken:
			t := tkn.Token()
			if t.String() == "</section>" && isPkgConstSection {
				return vals
			}
		case tt == html.TextToken:
			t := tkn.Token()
			if isPkgConstSection {
				vals = append(vals, t.Data)
			}
		}
	}
	return vals
}

func parsePkgConstFromSection(doc string) []string {
	constList := []string{}
	sectionItems := parseConstSection(doc)
	for _, i := range sectionItems {
		con := strings.TrimSpace(i)
		// Skip if comment
		if strings.HasPrefix(con, "//") {
			continue
		}
		if len(con) == 0 {
			continue
		}
		conKeySplit := strings.Split(con, "=")
		// Skip if is not in key=value format
		if len(conKeySplit) != 2 {
			continue
		}
		// Print const name
		constList = append(constList, conKeySplit[0])
	}
	return constList
}

func main() {
	content, err := getPackageDocContent()
	if err != nil {
		log.Fatal(err)
	}
	consts := parsePkgConstFromSection(string(content))
	// Print consts
	fmt.Println(`List of consts declared in "k8s.io/component-helpers/storage/volume":`)
	for _, c := range consts {
		fmt.Println(c)
	}
}
