package piscine

func IterativePower(nb int, power int) int {
	pwrs := 1

	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	for i := 1; i <= power; i++ {
		pwrs *= nb
	}

	return pwrs
}
