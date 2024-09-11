package main

import (
	"strings"
)

func digitMath (a string, b uint8, c int) string {
	result := ""
	if b == 42 {
		for i := 1; i <= c; i++ {
			result += a
		}
	}
	if b == 47 {
		result = a[:(len(a)/c)]
	}
	return result
}

func stringMath (a string, b uint8, c string) string {
	result := ""
	if b == 43 {
		result = a + c
	}
	if b == 45 {
		if len(c) > len(a) {
			panic("Некорректная операция")
		} else {
			result = strings.ReplaceAll(a, c, "")
		}
	}
	return result
}