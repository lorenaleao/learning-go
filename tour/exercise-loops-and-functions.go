package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	prev := 0.0
	count := 0
	for (z - prev) != 0 {
		//fmt.Println(prev, z)
		prev = z
		z -= (z*z - x) / (2 * z)
		//fmt.Println(count, z)
		count++
	}
	return z
}

func main() {
	var n float64 = 456
	sqrtn := Sqrt(n)
	fmt.Printf("Sqrt of %v: %v\n", n, sqrtn)
	fmt.Printf("Diff: %v\n", sqrtn - math.Sqrt(n))
}