package main

func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
}
