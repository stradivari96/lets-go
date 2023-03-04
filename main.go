package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
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