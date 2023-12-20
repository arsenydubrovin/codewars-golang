package main

func Summation(n int) int {
	// return (1 + n) * n / 2
	if n == 1 {
		return 1
	}
	return n + Summation(n-1)
}
