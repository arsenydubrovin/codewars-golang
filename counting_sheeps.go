package main

func CountSheeps(numbers []bool) (count int) {
	for _, n := range numbers {
		if n {
			count++
		}
	}
	return
}
