package piscine

func AppendRange(min, max int) []int {
	var answ []int

	for i := min - 1; i < max; i++ {
		if i != max-1 {
			answ = append(answ, i+1)
		}
	}
	return answ
}
