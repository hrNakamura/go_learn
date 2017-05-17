// 練習問題3.6　スーパーサンプリング
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const samplingCount = 2
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024 * samplingCount, 1024 * samplingCount
	)

	img := image.NewRGBA(image.Rect(0, 0, width/samplingCount, height/samplingCount))
	for py := 0; py < height; py += samplingCount {
		for px := 0; px < width; px += samplingCount {
			var val uint32
			for sx := 0; sx < samplingCount; sx++ {
				for sy := 0; sy < samplingCount; sy++ {
					y := float64(py+sy)/height*(ymax-ymin) + ymin
					x := float64(px+sx)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					v, _, _, _ := mandelbrot(z).RGBA()
					val += v
				}
			}
			val /= (samplingCount * samplingCount)
			col := color.RGBA{(uint8)(val), (uint8)(val), (uint8)(val), 255}
			// Image point (px, py) represents complex value z.
			img.Set(px/samplingCount, py/samplingCount, col)
		}

	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
