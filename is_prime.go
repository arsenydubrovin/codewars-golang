package main

import "math"

// t := time.Now()
// IsPrime(1234567891)
// fmt.Println(time.Since(t))
// 401µs

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
