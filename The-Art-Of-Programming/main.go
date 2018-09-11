//1.1 左旋转字符串

package main

import (
	"fmt"
)

func main() {
	data := LeftShiftOne("abcdef", 3)
	fmt.Println(data)

	s := []byte("abcdef")

	n := len(s)
	m := 3
	LeftRotateString(s, n, m)
	fmt.Println(string(s))
}

// 方法1
func LeftShiftOne(char string, n int) string {
	return string(char[n:len(char)]) + string(char[0:n])
}

// 三步反转法
func LeftRotateString(s []byte, m int, n int) {
	m %= n //若要左移动大于n位，那么和%n 是等价的
	ReverseString(s, 0, m-1)
	ReverseString(s, m, n-1)
	ReverseString(s, 0, n-1)
}

func ReverseString(s []byte, from int, to int) {
	for from < to {
		t := s[from]
		s[from] = s[to]
		s[to] = t
		from++
		to--
	}
}
