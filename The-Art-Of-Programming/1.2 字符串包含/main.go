// 字符串包含
/*

题目描述
给定两个分别由字母组成的字符串A和字符串B，字符串B的长度比字符串A短。请问，如何最快地判断字符串B中所有字母是否都在字符串A里？

为了简单起见，我们规定输入的字符串只包含大写英文字母，请实现函数bool StringContains(string &A, string &B)

比如，如果是下面两个字符串：

String 1：ABCD

String 2：BAD

答案是true，即String2里的字母在String1里也都有，或者说String2是String1的真子集。

如果是下面两个字符串：

String 1：ABCD

String 2：BCE

答案是false，因为字符串String2里的E字母不在字符串String1里。

同时，如果string1：ABCD，string 2：AA，同样返回true。
 */
package main

import (
	"fmt"
)

func main() {
	s1 := "ABCDEFG"
	s2 := "BCGD"
	fmt.Println(StringContain3(s1, s2))
}

// 1. 判断string2中的字符是否在string1中?最直观也是最简单的思路是，针对string2中每一个字符，逐个与string1中每个字符比较，看它是否在String1中。
func StringContain1(a string, b string) bool {
	for i := 0; i < len(b); i++ {
		var j int

		for j = 0; j < len(a); j++ {
			if a[j] == b[i] {
				break
			}
		}
		if (j >= len(a)) {
			return false
		}
	}
	return true
}

// 2. 如果允许排序的话，我们可以考虑下排序。比如可先对这两个字符串的字母进行排序，然后再同时对两个字串依次轮询。两个字串的排序需要(常规情况)O(m log m) + O(n log n)次操作，之后的线性扫描需要O(m+n)次操作。 关于排序方法，可采用最常用的快速排序，参考代码如下：
func StringContain2(a string, b string) bool {
	return true
}

/*
假设有一个仅由字母组成字串，让每个字母与一个素数对应，从2开始，往后类推，A对应2，B对应3，C对应5，......。遍历第一个字串，把每个字母对应素数相乘。最终会得到一个整数。

利用上面字母和素数的对应关系，对应第二个字符串中的字母，然后轮询，用每个字母对应的素数除前面得到的整数。如果结果有余数，说明结果为false。如果整个过程中没有余数，则说明第二个字符串是第一个的子集了（判断是不是真子集，可以比较两个字符串对应的素数乘积，若相等则不是真子集）。

思路总结如下：

按照从小到大的顺序，用26个素数分别与字符'A'到'Z'一一对应。
遍历长字符串，求得每个字符对应素数的乘积。
遍历短字符串，判断乘积能否被短字符串中的字符对应的素数整除。
输出结果。
如前所述，算法的时间复杂度为O(m+n)的最好的情况为O(n)（遍历短的字符串的第一个数，与长字符串素数的乘积相除，即出现余数，便可退出程序，返回false），n为长字串的长度，空间复杂度为O(1)。
 */
func StringContain3(a string, b string) bool {
	 p  := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,61, 67, 71, 73, 79, 83, 89, 97, 101}
	 f := 1
	 for i := 0; i < len(a); i++ {
	 	fmt.Println(a[i]-'A', p[a[i]-'A'])
		x := p[a[i]-'A']
		if (f % x != 0) {
			f *= x
		}
	}
	for i := 0; i < len(b); i++ {
		x := p[b[i]-'A']
		if f % x != 0 {
			return  false
		}
	}
	return  true
}
