package piscine

func ConcatParams(args []string) string {
	result := ""
	counter := 0
	for range args {
		counter++
	}
	for i := range args {
		if i == counter-1 {
			result += args[i]
		} else {
			result += args[i] + "\n"
		}
	}
	return result
}
