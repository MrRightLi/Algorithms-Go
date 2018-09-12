/*
最长回文子串
题目描述
给定一个字符串，求它的最长回文子串的长度。
 */

 /*
 分析与解法
最容易想到的办法是枚举所有的子串，分别判断其是否为回文。这个思路初看起来是正确的，但却做了很多无用功，如果一个长的子串包含另一个短一些的子串，那么对子串的回文判断其实是不需要的。

  */
package main

import "fmt"

func main() {
	fmt.Println(LongestPalindrome("heh"))
}

/*
解法一
那么如何高效的进行判断呢？我们想想，如果一段字符串是回文，那么以某个字符为中心的前缀和后缀都是相同的，例如以一段回文串“aba”为例，以b为中心，它的前缀和后缀都是相同的，都是a。

那么，我们是否可以可以枚举中心位置，然后再在该位置上用扩展法，记录并更新得到的最长的回文长度呢？答案是肯定的，参考代码如下：
 */
func LongestPalindrome(s string) int {
	var i, j, max, c int
	n := len(s)
	if (s == "" || n < 1) {
		return 0
	}
	max = 0
	for i = 0; i < n; i++ {
		for j = 0; (i - j >= 0) && (i + j < n); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			c = j * 2 + 1
		}
		if c > max {
			max = c
		}
		for j = 0; (i - j >= 0) && (i + j + 1 < n); j++  {
			if (s[i - j] != s[i + j + 1]) {
				break
			}
			c = j * 2 + 2
		}
		if c > max {
			max = c
		}
	}
	return max
}
