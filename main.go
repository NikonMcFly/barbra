package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pieperz/barbra/imageScaler"
)

// HealthzHandler ...
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running...\n"))
}

// ResizeHandler ...
func ResizeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Resizing")
	scale := imageScaler.NewTransformation()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&scale)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Allow user to upload there own image form the static site
	img, err := imageScaler.GetPng("./static/images/University of Houston Logo.png")
	if err != nil {
		log.Fatal(err)
	}

	scaledImg, err := imageScaler.NewScale(img, scale)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: [FUTURE] Send proccessed images to S3?
	out, err := os.Create("./static/images/resized_University of Houston Logo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, scaledImg)
	log.Print("...Done")
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/resize", ResizeHandler)
	r.HandleFunc("/healthz", HealthzHandler)

	// This will serve files under http://localhost:8000<filename>
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	srv := &http.Server{
		Handler: r, Addr: "127.0.0.1:8000",
		// Enforce timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("serving running on http://localhost:8000")
	log.Fatal(srv.ListenAndServe())
}
