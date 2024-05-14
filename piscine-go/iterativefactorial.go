package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	factr := 1

	if nb > 0 && nb <= 20 {
		for i := 1; i <= nb; i++ {
			factr *= i
		}
	} else if factr == 1 {
		return 0
	} else {
		return factr
	}
	return factr
}
