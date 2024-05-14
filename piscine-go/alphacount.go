package piscine

func AlphaCount(s string) int {
	count := 0
	for _, i := range s {
		if isAlpha(i) {
			count++
		}
	}
	return count
}

func isAlpha(alpha rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
