package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
	MeterPerFeet Meter = 0.3048
)

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FtToM convert Feet to Meter
func FtToM(f Feet) Meter { return Meter(f * Feet(MeterPerFeet)) }

// MToFt convert Meter to Feet
func MToFt(m Meter) Feet { return Feet(m / MeterPerFeet) }
