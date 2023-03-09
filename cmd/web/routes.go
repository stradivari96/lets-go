package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// To disable listing see:
	// https://www.alexedwards.net/blog/disable-http-fileserver-directory-listings
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Note: patterns ending in / are subtree path patterns
	// Order does not matter, priority by length
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
