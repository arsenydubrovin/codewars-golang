package main

import "strconv"

func Derive(coef, exp int) string {
	return strconv.Itoa(coef*exp) + "x^" + strconv.Itoa(exp-1)
}
