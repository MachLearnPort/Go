package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Get request (_ replaces err handling)
	bytes, _ := ioutil.ReadAll(resp.Body)                                        // reading the body of the response
	stringBody := string(bytes)                                                  // Convert to string
	fmt.Println(stringBody)
	resp.Body.Close() // Close the resources that made the requests
}
