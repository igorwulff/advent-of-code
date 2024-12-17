package utils

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func PowInt(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}

	return result
}
