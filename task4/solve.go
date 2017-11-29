package main

import (
	"unicode"
)

func RemoveEven(input []int) []int {
	result := make([]int, 0, 1)
	for _, val := range input {
		if (val % 2) != 0 {
			result = append(result, val)
		}
	}
	return result
}

func PowerGenerator(val int) func() int {
	prev := 1
	return func() (res int) {
		ret = prev * val
		prev = res
		return
	}
}

func DifferentWordsCount(str string) (count int) {
	var w string
	m := make(map[string]bool)
	for _, val := range str {
		if (val >= 'A' && val <= 'Z') || (val >= 'a' && val <= 'z')  {
			w += string(unicode.ToLower(val))
		} else {
			if len(w) != 0 {
				m[w] = true
			}
			w = ""
		}
	}
	if len(w) != 0 {
		m[w] = true
	}
	w = ""
	count = len(m)
	return
}
