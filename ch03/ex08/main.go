// 練習問題3.8　異なる数値計算によるマンデルブロ集合
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

type complexBigFloat struct {
	re *big.Float
	im *big.Float
}

type complexBigRat struct {
	re *big.Rat
	im *big.Rat
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	var fractal int
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "64":
			fractal = 1
		case "Float":
			fractal = 2
		case "Rat":
			fractal = 3
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if fractal == 1 {
		for py := 0; py < height; py++ {
			y := float32(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float32(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrotC64(z))
			}
		}

	} else {
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				switch fractal {
				case 2:
					z := complexBigFloat{new(big.Float).SetFloat64(x), new(big.Float).SetFloat64(y)}
					img.Set(px, py, mandelbrotBigFloat(z))
				case 3:
					z := complexBigRat{new(big.Rat).SetFloat64(x), new(big.Rat).SetFloat64(y)}
					img.Set(px, py, mandelbrotBigRat(z))
				default:
					z := complex(x, y)
					img.Set(px, py, mandelbrotC128(z))
				}
			}
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotC128(z complex128) color.Color {
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

func mandelbrotC64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		vAbs := math.Sqrt((float64)(real(v)*real(v) + imag(v)*imag(v)))
		if vAbs > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complexBigFloat) color.Color {
	const iterations = 10
	const contrast = 15
	th := new(big.Float).SetFloat64(2.0)

	var v complexBigFloat
	v = complexBigFloat{new(big.Float).SetFloat64(0), new(big.Float).SetFloat64(0)}
	for n := uint8(0); n < iterations; n++ {
		v = multpleComplexBigFloat(v, v)
		v = addComplexBigFloat(v, z)
		vAbs := absComplexBigFloat(v)
		if vAbs.Cmp(th) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z complexBigRat) color.Color {
	const iterations = 10
	const contrast = 15
	th := new(big.Rat).SetFloat64(2.0)

	var v complexBigRat
	v = complexBigRat{new(big.Rat).SetFloat64(0), new(big.Rat).SetFloat64(0)}
	for n := uint8(0); n < iterations; n++ {
		v = multpleComplexBigRat(v, v)
		v = addComplexBigRat(v, z)
		vAbs := absComplexBigRat(v)
		if vAbs.Cmp(th) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func multpleComplexBigFloat(x, y complexBigFloat) complexBigFloat {
	var z, tmp complexBigFloat
	z = complexBigFloat{new(big.Float).SetFloat64(0), new(big.Float).SetFloat64(0)}
	tmp = complexBigFloat{new(big.Float).SetFloat64(0), new(big.Float).SetFloat64(0)}
	//実部の計算
	z.re = z.re.Mul(x.re, y.re)
	tmp.re.Mul(x.im, y.im)
	z.re.Sub(z.re, tmp.re)
	//虚部の計算
	z.im.Mul(x.re, y.im)
	tmp.im.Mul(x.im, y.re)
	z.re.Add(z.im, tmp.im)
	return z
}

func multpleComplexBigRat(x, y complexBigRat) complexBigRat {
	var z, tmp complexBigRat
	z = complexBigRat{new(big.Rat).SetFloat64(0), new(big.Rat).SetFloat64(0)}
	tmp = complexBigRat{new(big.Rat).SetFloat64(0), new(big.Rat).SetFloat64(0)}
	//実部の計算
	z.re.Mul(x.re, y.re)
	tmp.re.Mul(x.im, y.im)
	z.re.Sub(z.re, tmp.re)
	//虚部の計算
	z.im.Mul(x.re, y.im)
	tmp.im.Mul(x.im, y.re)
	z.re.Add(z.im, tmp.im)
	return z
}

func addComplexBigFloat(x, y complexBigFloat) complexBigFloat {
	var z complexBigFloat
	z = complexBigFloat{new(big.Float).SetFloat64(0), new(big.Float).SetFloat64(0)}
	z.re.Add(x.re, y.re)
	z.im.Add(x.im, y.im)
	return z
}

func addComplexBigRat(x, y complexBigRat) complexBigRat {
	var z complexBigRat
	z = complexBigRat{new(big.Rat).SetFloat64(0), new(big.Rat).SetFloat64(0)}
	z.re.Add(x.re, y.re)
	z.im.Add(x.im, y.im)
	return z
}

func absComplexBigFloat(v complexBigFloat) *big.Float {
	r := v.re.Mul(v.re, v.re)
	i := v.im.Mul(v.im, v.im)
	r.Add(r, i)
	return sqrtBigFloat(r)
}

func absComplexBigRat(v complexBigRat) *big.Rat {
	r := v.re.Mul(v.re, v.re)
	i := v.im.Mul(v.im, v.im)
	r.Add(r, i)
	return sqrtBigRat(r)
}

func sqrtBigFloat(v *big.Float) *big.Float {
	const prec = 20
	// two := new(big.Float).SetPrec(prec).SetFloat64(2.0)
	steps := int(math.Log2(prec))
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	x := new(big.Float).SetPrec(prec).SetInt64(1)
	t := new(big.Float)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(v, x)    // t = 2.0 / x_n
		t.Add(x, t)    // t = x_n + (2.0 / x_n)
		x.Mul(half, t) // x_{n+1} = 0.5 * t
	}
	return x
}

func sqrtBigRat(v *big.Rat) *big.Rat {
	// two := new(big.Float).SetPrec(prec).SetFloat64(2.0)
	const prec = 10
	steps := int(math.Log2(prec))
	half := new(big.Rat).SetFloat64(0.5)
	x := new(big.Rat).SetInt64(1)
	t := new(big.Rat)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(v, x)    // t = 2.0 / x_n
		t.Add(x, t)    // t = x_n + (2.0 / x_n)
		x.Mul(half, t) // x_{n+1} = 0.5 * t
	}
	return x
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
