package main

import "fmt"

func main() {
	sum := 0
	
	// Unlike other languages like C, Java, or JavaScript 
	// there are no parentheses surrounding the three components 
	//of the for statement and the braces { } are always required.
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// The init and post statements are optional.
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// you can drop the semicolons: C's while is spelled for in Go.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// If you omit the loop condition it loops forever, so an 
	// infinite loop is compactly expressed.
	for {
	}
}