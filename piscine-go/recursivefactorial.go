package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb <= 0 {
		return 0
	} else if nb > 20 {
		return 0
	} else {
		return nb * IterativeFactorial(nb-1)
	}
}
