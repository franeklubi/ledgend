package ledgend


import (
    "math"
)


type Color struct {
    R, G, B uint8
}


// Gradient interpolates between two Colors at point m
//
// m represents a float64 between 0 and 1
func Gradient(a, b Color, m float64) (Color) {
    R := float64(a.R) - math.Floor( (float64(a.R)-float64(b.R))*m )
    G := float64(a.G) - math.Floor( (float64(a.G)-float64(b.G))*m )
    B := float64(a.B) - math.Floor( (float64(a.B)-float64(b.B))*m )

    return Color{uint8(R), uint8(G), uint8(B)}
}
