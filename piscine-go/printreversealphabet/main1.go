package main

import "github.com/01-edu/z01"

func main() {
	// making i loop from z to a
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	// print a new line
	z01.PrintRune('\n')
}
