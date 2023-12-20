package main

// stupid solution
func RepeatStr(repetitions int, value string) string {
	var res []rune
	for i := 0; i < repetitions; i++ {
		res = append(res, []rune(value)...)
	}
	return string(res)
}
