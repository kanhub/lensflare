package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// Fresh continues runnign the last version of the errorless build
	// pay attention to the logs in Fresh
	// fresh has default config, but you can have your own config so it watches
	// files that change and the reload based on files you care about
	w.Header().Set("Content-Type", "text/html")
	fmt.Println("Someone visited our page")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome my website</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<p>To get in touch, please send an email to <a href=\"mailto:support@lensflare.com\">support@lensflare.com</a></p>")
	} else {
		// WriteHeader to grab the 404 status constant within the http package
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Could not find that page :(</h1>")
	}
}

func main() {
	// built in servemux
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
