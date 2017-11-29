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
	return func() (ret int) {
		ret = prev * val
		prev = ret
		return
	}
}

func DifferentWordsCount(str string) (cnt int) {
	var word string
	m := make(map[string]bool)
	for _, val := range str {
		if (val >= 'a' && val <= 'z') || (val >= 'A' && val <= 'Z') {
			word += string(unicode.ToLower(val))
		} else {
			if len(word) != 0 {
				m[word] = true
			}
			word = ""
		}
	}
	if len(word) != 0 {
		m[word] = true
	}
	word = ""
	cnt = len(m)
	return
}
