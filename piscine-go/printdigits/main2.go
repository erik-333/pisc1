package main

import "github.com/01-edu/z01"

func main() {
	// making i loop from 0 to 9
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	// print a new line
	z01.PrintRune('\n')
}
