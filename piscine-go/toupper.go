package piscine

func ToUpper(s string) string {
	runeArray := []rune(s)
	for t := range runeArray {
		if runeArray[t] >= 'a' && runeArray[t] <= 'z' {
			runeArray[t] = runeArray[t] - 32
		}
	}
	return string(runeArray)
}
