package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	// abs()
	// exp()
	polar()
}

func polar() {
	r, theta := cmplx.Polar(2i)
	fmt.Printf("r: %.1f, θ: %.1f*π", r, theta/math.Pi)
}

func exp() {
	fmt.Printf("%.1f", cmplx.Exp(1i*math.Pi)+1)
}

func abs() {
	fmt.Printf("%.1f", cmplx.Abs(3+4i))
}
