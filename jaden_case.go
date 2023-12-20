package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToJadenCase(str string) string {
	caser := cases.Title(language.English)
	return caser.String(str)
}
