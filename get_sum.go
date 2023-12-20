package main

func GetSum(a, b int) (sum int) {
	if a > b {
		a, b = b, a
	}
	return (a + b) * (b - a + 1) / 2
}
