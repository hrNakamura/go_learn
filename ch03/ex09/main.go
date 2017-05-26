// 練習問題3.9　クライアントにマンデルブロ集合を出力するサーバ
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var offsetX float64
var offsetY float64
var iterate uint8

var regX = regexp.MustCompile(`x=\d`)
var regY = regexp.MustCompile(`y=\d`)
var regN = regexp.MustCompile(`n=\d`)

//TODO X,Yはフラクタル画像の中心点、倍率は画像の拡大縮小率

func main() {
	http.HandleFunc("/", route) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func route(w http.ResponseWriter, r *http.Request) {
	if regX.MatchString(r.URL.Path) {
		offsetX = getParam(regX.FindString(r.URL.Path))
	}
	if regY.MatchString(r.URL.Path) {
		offsetY = getParam(regY.FindString(r.URL.Path))
	}
	if regN.MatchString(r.URL.Path) {
		iterate = getIterations(regN.FindString(r.URL.Path))
	}
	w.Header().Set("Content-type", "image/png")
	w.WriteHeader(http.StatusOK)
	writeMandelbrot(w)
	w.(http.Flusher).Flush()
}

func getParam(str string) float64 {
	strs := strings.Split(str, "=")
	var param float64
	var err error
	param, err = strconv.ParseFloat(strs[len(strs)-1], 64)
	if err != nil {
		fmt.Printf("%v,\n", err)
	}
	return param
}

func getIterations(str string) uint8 {
	strs := strings.Split(str, "=")
	var param uint64
	var err error
	param, err = strconv.ParseUint(strs[len(strs)-1], 10, 8)
	if err != nil {
		fmt.Printf("%v,\n", err)
	}
	return uint8(param)
}

func writeMandelbrot(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	var iterations uint8
	const contrast = 15
	if iterate != uint8(0) {
		iterations = iterate
	} else {
		iterations = 200
	}
	iterations = 100

	v := complex(offsetX, offsetY)
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
