package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup()

type siteMapIndex struct {
	Locations []string `xml:"sitemap>loc"` // []string is a slice (i.e. [5]type is an array of fixed size (where type is int, float, etc), where as []type is a slice
	// []Location is a slice of type location. Also, remeber Locations and Location need to have the first letter CAPS to make public
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s siteMapIndex
	var n News

	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Get request (_ replaces err handling)
	bytes, _ := ioutil.ReadAll(resp.Body)                                        // reading the body of the response
	resp.Body.Close()                                                            // Close the resources that made the requests
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	// fmt.Println(s.Locations)

	for _, Location := range s.Locations { // _ becuase we are not using that placeholder. Range iterates over our struture
		resp, _ := http.Get(Location) // Get request (_ replaces err handling)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n) // reading the body of the response
		resp.Body.Close()
		//fmt.Println(n.Keywords)
		
		//Creating out MAP with titles as our key
		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
