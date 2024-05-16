package piscine

func IsLower(s string) bool {
	for _, l := range s {
		if l >= 'a' && l <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
