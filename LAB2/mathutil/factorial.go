package mathutil

func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	result := 1
	for value := 2; value <= n; value++ {
		result *= value
	}
	return result
}
