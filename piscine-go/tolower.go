package piscine

func ToLower(s string) string {
	runeArray := []rune(s)
	for l := range runeArray {
		if runeArray[l] >= 'A' && runeArray[l] <= 'Z' {
			runeArray[l] = runeArray[l] + 32
		}
	}
	return string(runeArray)
}
