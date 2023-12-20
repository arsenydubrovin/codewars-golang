package main

func ReverseSeq(n int) (ret []int) {
	for n > 0 {
		ret = append(ret, n)
		n--
	}
	return
}
