package main

import (
	"encoding/json"
	"flag"
	"image/png"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/pieperz/barbra/imageScale"
	"golang.org/x/exp/shiny/unit"
)

type point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type line struct {
	Start *point `json:"start"`
	End   *point `json:"end"`
}

// Scale ...
type Scale struct {
	Line   *line   `json:"line"`
	Length float64 `json:"length"`
}

func newPoint() *point {
	return &point{
		X: 0,
		Y: 0,
	}
}

func newLine() *line {
	return &line{
		Start: newPoint(),
		End:   newPoint(),
	}
}

// NewScale ...
func NewScale() *Scale {
	return &Scale{
		Line:   newLine(),
		Length: 0,
	}
}

// HealthzHandler ...
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running...\n"))
}

// ResizeHandler ...
func ResizeHandler(w http.ResponseWriter, req *http.Request) {

	scale := NewScale()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&scale)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Allow user to upload there own image form the static site
	img, _ := imageScale.GetPng("./static/images/University of Houston Logo.png")

	// TODO: Add this to the strut as methods
	pixelScale := unit.Pixels(float64(scale.Line.End.X - scale.Line.Start.X))
	knownLength := unit.Inches(float64(scale.Length))
	// TODO: add Axis

	// TODO just pass in a Scale Oject
	scaledImg, err := imageScale.ScaleImage(img, pixelScale, knownLength, "x")
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
