package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb == 2 {
		return 0
	}
	if nb > 0 {
		rslt := 1
		sqrt := 0
		for a := 1; rslt <= nb; a++ {
			rslt = a * a
			sqrt++
			if rslt == nb {
				return sqrt
			}
		}
		return 0
	} else {
		return 0
	}
}
