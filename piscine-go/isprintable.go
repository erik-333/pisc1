package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)

	for _, p := range runes {
		if !(p >= 32 && p <= 127) {
			return false
		}
	}
	return true
}
