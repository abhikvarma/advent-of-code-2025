package utils

func FloorDiv(a, b int) int {
	if a >= 0 {
		return a / b
	}

	// -10/100 -> -1
	// -100/100 -> -1
	return (a - b + 1) / b
}
