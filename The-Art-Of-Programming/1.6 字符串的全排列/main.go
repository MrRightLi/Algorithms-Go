/*
题目描述
输入一个字符串，打印出该字符串中字符的所有排列。

例如输入字符串abc，则输出由字符a、b、c 所能排列出来的所有字符串

abc、acb、bac、bca、cab 和 cba。
 */

package main

func main() {

}

/*
解法一、递归实现
从集合中依次选出每一个元素，作为排列的第一个元素，然后对剩余的元素进行全排列，如此递归处理，从而得到所有元素的全排列。以对字符串abc进行全排列为例，我们可以这么做：以abc为例

固定a，求后面bc的排列：abc，acb，求好后，a和b交换，得到bac
固定b，求后面ac的排列：bac，bca，求好后，c放到第一位置，得到cba
固定c，求后面ba的排列：cba，cab。
 */

func CalcAllPermutation(str string) []string {
	var stringRank []string
	for i := 0; i < len(str); i++ {
		stringTem  := string(str[i])
		for j := 0; j < len(str); j++ {
			if i != j {
				stringTem = stringTem + string(str[j])
			}
		}
		stringRank = append(stringRank, string(stringTem))
	}
	return stringRank
}
