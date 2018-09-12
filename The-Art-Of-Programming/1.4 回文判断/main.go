package main

import (
	"fmt"
)

func main() {
	str := "helsaleh"

	fmt.Println(IsPalindrome(str))
}

func IsPalindrome(s string) bool {
	// 非法输入
	n := len(s)
	if s == "" || n < 1 {
		return false
	}

	front := 0
	back := len(s) - 1

	for front < back  {
		if s[front] != s[back] {
			return false
		}
		front++
		back--
	}
	return true
}