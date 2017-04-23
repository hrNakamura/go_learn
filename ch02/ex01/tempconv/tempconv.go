// Package tempconv performs Celsius and Fahrenheit conversions.
// add Kelvin
package tempconv

import "fmt"

// Celsius is a celsius temperture type name
type Celsius float64

// Fahrenheit is a Fahrenheit temperture type name
type Fahrenheit float64

// Kelvin is a Kelvin temperture type name
type Kelvin float64

// tempature const value
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

//!-
