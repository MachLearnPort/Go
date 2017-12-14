package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func main() {
	var s siteMapIndex
	var n News

	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Get request (_ replaces err handling)
	bytes, _ := ioutil.ReadAll(resp.Body)                                        // reading the body of the response
	resp.Body.Close()                                                            // Close the resources that made the requests
	xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)

	for _, Location := range s.Locations { // _ becuase we are not using that placeholder. Range iterates over our struture
		resp, _ := http.Get(Location) // Get request (_ replaces err handling)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n) // reading the body of the response
		//fmt.Println(n.Keywords)
		// resp.Body.Close()

		//Creating out MAP with titles as our key
		for idx := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
