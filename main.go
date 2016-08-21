package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// HealthzHandler ...
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running...\n"))
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/healthz", HealthzHandler)

	// This will serve files under http://localhost:8000<filename>
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Enforce timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("serving files on http://localhost:8000")
	log.Fatal(srv.ListenAndServe())
}
