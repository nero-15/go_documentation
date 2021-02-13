package main

import (
	"fmt"
	"math"
)

func main() {
	// abs()
	// acos()
	// acosh()
	// cbrt()
	// ceil()
	// copysign()
	// cos()
	// dim()
	// exp()
	// exp2()
	// expm1()
	// floor()
	// log10()
	// mod()
	// pow()
	// round()
	// sqrt()
	trunc()
}

func trunc() {
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))
	fmt.Printf("%.2f\n", math.Trunc(-1.2345))
}

func sqrt() {
	const (
		a = 3
		b = 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%.1f", c)
}

func round() {
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)
}

func pow() {
	c := math.Pow(2, 3)
	fmt.Printf("%.1f", c)
}

func mod() {
	c := math.Mod(7, 4)
	fmt.Printf("%.1f", c)
}

func log10() {
	fmt.Printf("%.1f", math.Log10(100))
}

func floor() {
	c := math.Floor(1.51)
	fmt.Printf("%.1f", c)
}

func expm1() {
	fmt.Printf("%.6f\n", math.Expm1(0.01))
	fmt.Printf("%.6f\n", math.Expm1(-1))
}

func exp2() {
	fmt.Printf("%.2f\n", math.Exp2(1))
	fmt.Printf("%.2f\n", math.Exp2(-3))
}

func exp() {
	fmt.Printf("%.2f\n", math.Exp(1))
	fmt.Printf("%.2f\n", math.Exp(2))
	fmt.Printf("%.2f\n", math.Exp(-1))
}

func dim() {
	fmt.Printf("%.2f\n", math.Dim(4, -2))
	fmt.Printf("%.2f\n", math.Dim(-4, 2))
}

func cos() {
	fmt.Printf("%.2f", math.Cos(math.Pi/2))
}

func copysign() {
	fmt.Printf("%.2f", math.Copysign(3.2, -1))
}

func ceil() {
	c := math.Ceil(1.49)
	fmt.Printf("%.1f", c)
}

func cbrt() {
	fmt.Printf("%.2f\n", math.Cbrt(8))
	fmt.Printf("%.2f\n", math.Cbrt(27))
}

func acosh() {
	fmt.Printf("%.2f", math.Acosh(1))
}

func acos() {
	fmt.Printf("%.2f", math.Acos(1))
}

func abs() {
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	y := math.Abs(2)
	fmt.Printf("%.1f\n", y)
}
