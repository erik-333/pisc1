package piscine

func StrRev(s string) string {
	rev := ""
	for _, a := range s {
		rev = string(a) + rev
	}
	return rev
}
