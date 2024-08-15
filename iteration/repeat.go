package iteration

import "strings"

const repeatCount = 5

func Repeat(char string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}

func Repeat2(char string) string {
	return strings.Repeat(char, repeatCount)
}
