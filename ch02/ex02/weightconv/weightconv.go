package weightconv

import "fmt"

type Pond float64
type Kilogram float64

const (
	KilogramPerPond Kilogram = 0.4536
)

func (p Pond) String() string     { return fmt.Sprintf("%glb", p) }
func (g Kilogram) String() string { return fmt.Sprintf("%gkg", g) }

// PToKg convert Pond to Kilogram
func PToKg(p Pond) Kilogram { return Kilogram(p * Pond(KilogramPerPond)) }

// KgToP convert Kilogram to Pond
func KgToP(g Kilogram) Pond { return Pond(g / KilogramPerPond) }
