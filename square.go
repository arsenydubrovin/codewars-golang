package main

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 0; i < 10; i++ {
		shift := (z*z - x) / (2 * z)
		z -= shift
		if shift < 0.00001 {
			break
		}
	}
	return z
}
