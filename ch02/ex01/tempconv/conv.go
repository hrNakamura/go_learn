package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CtoK converts a Celsius temperture to Kelvin
func CtoK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KtoC converts a Kelvin temperture to Celsius
func KtoC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

//!-
