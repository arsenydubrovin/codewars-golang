package main

func Factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}
