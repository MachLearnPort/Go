package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there</h1>")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>... and simple</p>")
	fmt.Fprintf(w, "<p>You can %s even add %s</p>", "can", "<strong>variables</strong>")
	fmt.Fprintf(w, `<h1> This </h1>
									<h1> is </h1>
									<h1> Multiline </h1>`)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
