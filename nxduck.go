package nxduck

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GenerateSearchURL creates a DuckDuckGo URL for searching
func GenerateSearchURL(searchTerm string) string {
	searchTerm = url.QueryEscape(searchTerm)
	return "https://duckduckgo.com/html/?q=" + searchTerm
}

/*
// GetSearchResults gets Search results from URL
func GetSearchResults(duckURL string) []string {
	var urlsFound []string
	//get webpage html
	doc, err := goquery.NewDocument(duckURL)
	if err != nil {
		log.Fatal(err)
	}
	//find "links_main"
	sel := doc.Find(".links_main")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		// find "a"+
		found := single.Find("a")
		// find a's href
		itemURL, _ := found.Attr("href")

		urlsFound = append(urlsFound, itemURL)
	}
	return urlsFound
}
*/

//FIXME these functions are a mess and need to be rationalized

// GetSearchResultURLs returns URL search results
func GetSearchResultURLs(duckURL string) []string {
	var urlsFound []string
	//get webpage html
	doc, err := goquery.NewDocument(duckURL)
	if err != nil {
		log.Fatal(err)
	}
	//find "links_main"
	sel := doc.Find(".links_main")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		// find "a"+
		found := single.Find("a")
		// find a's href
		itemURL, _ := found.Attr("href")

		urlsFound = append(urlsFound, itemURL)
	}
	return urlsFound
}

// GetSearchResultObjects returns URL and Title of search results via a SearchResult object
func GetSearchResultObjects(duckURL string) []SearchResult {
	//var urlsFound [][]string
	var results []SearchResult
	//get webpage html
	doc, err := goquery.NewDocument(duckURL)
	if err != nil {
		log.Fatal(err)
	}
	//find "links_main"
	sel := doc.Find(".links_main")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		// find "a"+
		found := single.Find("a")
		// find a's href
		itemURL, _ := found.Attr("href")

		var tmpItem SearchResult

		mayTitles := single.Find(".result__title")
		for j := range mayTitles.Nodes {
			mayTitle := mayTitles.Eq(j)
			//fmt.Printf(">>>>>%s<<<<<\n", mayTitle.Text())
			tmpItem.Title = strings.TrimSpace(mayTitle.Text())
			if strings.Contains(tmpItem.Title, "...") {
				tmpItem.IncompleteTitle = true
			}
			//fmt.Printf("(%t)>>>>>%s<<<<<\n", tmpItem.IncompleteTitle, tmpItem.Title)
			break

		}
		//tmpItem.Title = found.Text()
		tmpItem.URL = itemURL

		results = append(results, tmpItem)

		//urlsFound = append(urlsFound, tmpItem)
	}
	return results
}

/*
//GetSearchResults returns URL and Title of search results
func GetSearchResults(duckURL string) [][]string {
	var results [][]string
	var Titles []string
	var URLs []string
	//get webpage html
	doc, err := goquery.NewDocument(duckURL)
	if err != nil {
		log.Fatal(err)
	}
	//find "links_main"
	sel := doc.Find(".links_main")
	for i := range sel.Nodes {
		single := sel.Eq(i)
		// find "a"+
		found := single.Find("a")
		// find a's href
		itemURL, _ := found.Attr("href")

		//var tmpItem SearchResult

		//tmpItem.Title = found.Text()
		//tmpItem.URL = itemURL

		Titles = append(Titles, found.Text())
		URLs = append(URLs, itemURL)

		//urlsFound = append(urlsFound, tmpItem)
	}
	results = append(results, URLs)
	results = append(results, Titles)
	return results
}
*/
