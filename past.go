package main

func Past(h, m, s int) int {
	return h*3600000 + 60000*m + s*1000
}
