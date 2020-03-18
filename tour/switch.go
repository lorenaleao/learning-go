package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	// It runs the first case whose value is equal to the condition expression.
	// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except
	// that Go only runs the selected case, not all the cases that follow. In 
	// effect, the break statement that is needed at the end of each case in 
	// those languages is provided automatically in Go. Another important 
	// difference is that Go's switch cases need not be constants, and the values
	//  involved need not be integers.

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}