package main

import "fmt"

// 无重复字符的最长子串
/* example
** input: s = "abcabcbb"
** output: 3
** explain: 因为无重复字符的最长子串是"abc",所以其长度为3
 */

func lenghtOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]] ++
			rk++
		}
		ans = max(ans, rk - i + 1)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main()  {
	s := "abcabcbb"
	fmt.Println(lenghtOfLongestSubstring(s))
}
