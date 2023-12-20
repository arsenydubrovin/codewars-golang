package main

func Between(a, b int) []int {
	var ret []int
	for i := a; i <= b; i++ {
		ret = append(ret, i)
	}
	return ret
}
