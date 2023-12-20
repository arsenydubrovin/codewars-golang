package main

import "fmt"

func Digitize(n int) (result []int) {
	str := fmt.Sprintf("%d", n)

	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, int(str[i])-48)
	}
	return
}
