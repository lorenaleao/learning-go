package main

import "fmt"

//doesntwork := true
// you get a syntax error: non-declaration statement outside function body

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}