// 練習問題3.3　高さに応じた点の色付け
package main

import (
	"fmt"
	"math"
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

type Points struct {
	z, ax, ay, bx, by, cx, cy, dx, dy float64
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	pList := make([]Points, 0, cells*cells)
	var zMax, zMin float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ps Points
			// 有限でない結果の場合はその点は描画せずにスキップ
			ps.ax, ps.ay, ps.z = corner(i+1, j)
			if math.IsInf(ps.z, 0) || math.IsNaN(ps.z) {
				fmt.Printf("<!-- corner(%v, %v) is not finite -->\n", i, j)
				continue
			}
			ps.bx, ps.by, _ = corner(i, j)
			ps.cx, ps.cy, _ = corner(i, j+1)
			ps.dx, ps.dy, _ = corner(i+1, j+1)

			if ps.z > zMax {
				zMax = ps.z
			}
			if ps.z < zMin {
				zMin = ps.z
			}
			pList = append(pList, ps)
		}
	}
	for _, ps := range pList {
		v := (ps.z - zMin) / (zMax - zMin)
		r, g, b := colorScale(v)
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke=\"#%02X%02X%02X\"/>\n",
			ps.ax, ps.ay, ps.bx, ps.by, ps.cx, ps.cy, ps.dx, ps.dy, r, g, b)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//参考http://takacity.blog.fc2.com/blog-entry-69.html
func colorScale(v float64) (uint8, uint8, uint8) {
	var r, g, b uint8
	if v < 0.25 {
		b = 255
		g = (uint8)(255.0 * math.Sin(v*2.0*math.Pi))
		r = 0
	} else if v < 0.5 {
		b = (uint8)(255.0 * math.Sin(v*2.0*math.Pi))
		g = 255
		r = 0
	} else if v < 0.75 {
		b = 0
		g = 255
		r = (uint8)(-255.0 * math.Sin(v*2.0*math.Pi))
	} else {
		b = 0
		g = (uint8)(-255.0 * math.Sin(v*2.0*math.Pi))
		r = 255
	}
	return r, g, b
}

//!-
