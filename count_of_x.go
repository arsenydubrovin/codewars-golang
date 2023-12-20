package main

func CountBy(x, n int) []int {
	number := 0
	result := []int{}
	for i := 0; i < n; i++ {
		number += x
		result = append(result, number)
	}
	return result
}
