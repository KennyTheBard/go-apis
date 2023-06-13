package services

func Factorial(n int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Factorial(n-1) + Factorial(n-2)
}
