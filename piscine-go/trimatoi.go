package piscine

func TrimAtoi(s string) int {
	result := 0
	ngtv := false
	foundNum := false
	for _, char := range s {
		if !foundNum && char == '-' {
			ngtv = true
		}
		if char >= '0' && char <= '9' {
			foundNum = true
			result = result*10 + int(char) - 48
		}
	}
	if ngtv {
		return -result
	}
	return result
}
