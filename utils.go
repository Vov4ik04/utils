package utils

func Contains(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return true
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return true
}
