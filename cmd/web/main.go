package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

// for easier dependency injection
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var cfg config

func main() {
	// Go prefers command line flags, type convertion & help
	// go run ./cmd/web -addr=$SNIPPETBOX_ADDR
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// initialize a http.Server struct to change the ErrorLog
	// as with the default http.ListenAndServe we cant
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.addr)

	// avoid DefaultServeMux, http.ListenAndServe(":4000", nil)
	// any package could add a handle there
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
