package main

func PowersOfTwo(n int) (ret []uint64) {
	for i := 0; i <= n; i++ {
		ret = append(ret, 1<<i)
	}
	return
}
