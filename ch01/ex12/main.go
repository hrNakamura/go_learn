// ex12 is lissajous gif animation server.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var cycles int  // number of complete x oscillator revolutions
var res float64 // angular resolution
var size int    // image canvas covers [-size..+size]
var nframes int // number of animation frames
var delay int   // delay between frames in 10ms units

func main() {
	cycles = 5
	res = 0.001
	size = 100
	nframes = 64
	delay = 8
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/cycles=", cycleParam)
	http.HandleFunc("/res=", resParam)
	http.HandleFunc("/size=", sizeParam)
	http.HandleFunc("/nframes=", nframeParam)
	http.HandleFunc("/delay=", delayParam)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func cycleParam(w http.ResponseWriter, r *http.Request) {
	strlist := strings.Split(r.URL.Path, "=")
	if len(strlist) == 2 {
		if p, err := strconv.Atoi(strlist[1]); err != nil {
			cycles = p
			lissajous(w)
		}
	}
}

func resParam(w http.ResponseWriter, r *http.Request) {
	strlist := strings.Split(r.URL.Path, "=")
	if len(strlist) == 2 {
		if p, err := strconv.ParseFloat(strlist[1], 64); err != nil {
			res = p
			lissajous(w)
		}
	}
}

func sizeParam(w http.ResponseWriter, r *http.Request) {
	strlist := strings.Split(r.URL.Path, "=")
	if len(strlist) == 2 {
		if p, err := strconv.Atoi(strlist[1]); err != nil {
			size = p
			lissajous(w)
		}
	}
}

func nframeParam(w http.ResponseWriter, r *http.Request) {
	strlist := strings.Split(r.URL.Path, "=")
	if len(strlist) == 2 {
		if p, err := strconv.Atoi(strlist[1]); err != nil {
			nframes = p
			lissajous(w)
		}
	}
}

func delayParam(w http.ResponseWriter, r *http.Request) {
	strlist := strings.Split(r.URL.Path, "=")
	if len(strlist) == 2 {
		if p, err := strconv.Atoi(strlist[1]); err != nil {
			delay = p
			lissajous(w)
		}
	}
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-
