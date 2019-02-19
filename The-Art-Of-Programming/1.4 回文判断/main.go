package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "A man, a plan, a canal: Panama"
	//str := ""
	str := "0P"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {
	// 非法输入
	n := len(s)
	if s == "" || n < 1 {
		return true
	}
	s = strings.ToLower(s)
	front := 0
	back := len(s) - 1

	for front < back  {
		for front < back {
			if (s[front] < 'a' || s[front] > 'z') && (s[front] < '0' || s[front] > '9'){
				front++
				continue
			}
			break
		}
		for front < back {
			if (s[back] < 'a' || s[back] > 'z') && (s[back] < '0' || s[back] > '9'){
				back--
				continue
			}
			break
		}
		if s[front] != s[back] {
			return false
		}
		front++
		back--
	}
	return true
}