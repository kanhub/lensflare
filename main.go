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
	fmt.Println("Someone visited our page")
	fmt.Fprint(w, "<h1>Welcome my website</h1>")
}

func main() {
	// built in servemux
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
