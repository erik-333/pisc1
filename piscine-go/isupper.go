package piscine

func IsUpper(s string) bool {
	for _, u := range s {
		if u >= 'A' && u <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
