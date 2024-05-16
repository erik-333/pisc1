package piscine

func IsAlpha(s string) bool {
	for _, a := range s {
		if a >= 'a' && a <= 'z' {
		} else if a >= 'A' && a <= 'Z' {
		} else if a >= '0' && a <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
