package main

import (
	"fmt"
)

/**
 * 环绕字符串中唯一的子字符串
 * 把字符串s看作是"abcdefghijklmnopqrstuvwxyz
的无限环绕字符串，所以s看起来是这样的：
	"...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...."
现在给定另一个字符串p,返回s中惟一的p的非空子串的数量。
  example1:
		input:  p="zab"
		output: 6
        explain: 字符串s中有6个子串"z,a,b,za,ab,zab"
*/

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func findSubstringInWraproundString(p string) int {
	i := 0
	cnum := make([]int, 26)
	for i < len(p) {
		start, curLen := int(p[i]-'a'), 1
		i++
		for i < len(p) && ((p[i] == 97 && p[i-1] == 122) || p[i]-p[i-1] == 1) {
			curLen++
			i++
		}
		for j := 0; j < min(curLen, 26); j++ {
			if cnum[(start + j) %26] < curLen - j {
				cnum[(start + j) %26] = curLen - j
			}
		}
	}
	result := 0
	for _, v := range cnum {
		result += v
	}
	return result
}

func main() {
	var s string
	all := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 1000; i++ {
		for j := 0; j < len(all); j++ {
			s += string(all[j])
		}
	}

	result := findSubstringInWraproundString("cac")
	fmt.Println(result)
	fmt.Println('z' - 'a')
}
