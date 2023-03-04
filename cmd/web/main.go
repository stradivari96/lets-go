package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// To disable listing see:
	// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Note: patterns ending in / are subtree path patterns
	// Order do not matter, longest ones have priority
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on :4000")
	// you can also use the DefaultServeMux via
	// http.HandleFunc and http.ListenAndServe(":4000", nil)
	// but any package could add a handle there
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
