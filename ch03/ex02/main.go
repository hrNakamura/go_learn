// 練習問題3.2　ほかのmathパッケージ関数によるその他の図形の可視化
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var fType int64

func main() {
	if len(os.Args) > 1 {
		v, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err == nil {
			fType = v
		}
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// 有限でない結果の場合はその点は描画せずにスキップ
			ax, ay, isOK := corner(i+1, j)
			if !isOK {
				fmt.Printf("<!-- corner(%v, %v) is not finite -->\n", i+1, j)
				continue
			}
			bx, by, isOK := corner(i, j)
			if !isOK {
				fmt.Printf("<!-- corner(%v, %v) is not finite -->\n", i, j)
				continue
			}
			cx, cy, isOK := corner(i, j+1)
			if !isOK {
				fmt.Printf("<!-- corner(%v, %v) is not finite -->\n", i, j+1)
				continue
			}
			dx, dy, isOK := corner(i+1, j+1)
			if !isOK {
				fmt.Printf("<!-- corner(%v, %v) is not finite -->\n", i+1, j+1)
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	// z := f(x, y)
	var z float64
	switch fType {
	case 1:
		z = f2(x, y)
	case 2:
		z = f3(x, y)
	case 3:
		z = f4(x, y)
	default:
		z = f(x, y)
	}
	// 無限大またはNaNであればfalseを返す
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func f2(x, y float64) float64 {
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12.0
}

func f3(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(-x) * math.Pow(1.5, -r)
}

func f4(x, y float64) float64 {
	return math.Sin(x*y/10.0) / 10.0
}

//!-
