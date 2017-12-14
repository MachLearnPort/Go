package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type siteMapIndex struct {
	Locations []Location `xml:"sitemap"` // []Location is a slice (i.e. [5]type is an array of fixed size (where type is int, float, etc), where as []type is a slice
	// []Location is a slice of type location. Also, remeber Locations and Location need to have the first letter CAPS to make public
}

type Location struct {
	Loc string `xml:"loc"` // Makes Loc the data under the <loc> tag in the XML data
}

// value receiver, since we are just extracting data and not trying to modify it
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Get request (_ replaces err handling)
	bytes, _ := ioutil.ReadAll(resp.Body)                                        // reading the body of the response
	resp.Body.Close()                                                            // Close the resources that made the requests

	var s siteMapIndex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	for _, Location := range s.Locations { // _ becuase we are not using that placeholder
		fmt.Printf("\n%s", Location)
	}
}
