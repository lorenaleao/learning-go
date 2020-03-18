package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	//test: 
	//var z uint = f 
	//cannot use f (type float64) as type uint in assignment
	fmt.Println(x, y, z)
}