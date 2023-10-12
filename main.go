package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var namespaces = []string{
	"site",
	"text",
	"title",
	"intext",
	"date",
}

type Terms struct {
	PositiveTerms []string
	NevativeTerms []string
}

func getNamespacesField(query string) (map[string]string, string) {

	join := strings.Join(namespaces, `|`)
	pattern := regexp.MustCompile(`(` + join + `):\s?("[^"]+"|\S+)`)

	var result = make(map[string]string)

	matches := pattern.FindAllStringSubmatch(query, -1)

	query = pattern.ReplaceAllString(query, "")
	for _, match := range matches {
    result[capitalize(match[1])] = match[2]
	}

	return result, query
}

func getTerms(query string) (Terms, string) {

	pattern := regexp.MustCompile(`-?("[^"]+"|\S+)`)

	matches := pattern.FindAllStringSubmatch(query, -1)

	var terms Terms
	// fmt.Println(query)
	for _, match := range matches {
		if strings.HasPrefix(match[0], "-") {
			terms.NevativeTerms = append(terms.NevativeTerms, match[0])
		} else {
			terms.PositiveTerms = append(terms.PositiveTerms, match[0])
		}
	}

	pattern.ReplaceAllString(query, "")
	return terms, query
}

func capitalize(text string) string {

	runes := []rune(text)
	runes[0] = unicode.ToUpper(runes[0])

	text = string(runes)

	return text
}

func main() {
	query := `site:github.com intext:"python language" date:>03300d random -term -"quoted term 2" asdasd das `

	fields, newQuery := getNamespacesField(query)
	terms, newQuery := getTerms(newQuery)

	for i, term := range terms.NevativeTerms {
		fmt.Printf("Negative Term %d: %s\n", i+1, term)
	}

	for i, term := range terms.PositiveTerms {
		fmt.Printf("Positive Term %d: %s\n", i+1, term)
	}

  for k, v := range fields {
    fmt.Printf("fieldSearch %s: %s\n", k, v)
  }

}
