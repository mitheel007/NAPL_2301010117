package mathutil

func Power(base, exponent int) int {
	if exponent < 0 {
		return 0
	}

	result := 1
	for value := 0; value < exponent; value++ {
		result *= base
	}
	return result
}
