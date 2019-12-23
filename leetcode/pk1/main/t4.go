package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func longestPalindrome(s string) string {
	var start, end int

	if s == "" {
		return ""
	} else if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		len1 := center(s, i, i)
		len2 := center(s, i, i+1)
		len3 := 0
		if len1 > len2 {
			len3 = len1
		} else {
			len3 = len2
		}

		if len3 > end-start {
			start = i - (len3-1)/2
			end = i + len3/2
		}
	}

	return s[start : end+1]
}

func center(s string, left, right int) int {
	l := left
	r := right

	for l >= 0 && r < len(s) && s[l:l+1] == s[r:r+1] {
		l--
		r++
	}
	return r - l - 1
}

func main() {
	s := "a"
	fmt.Println(longestPalindrome(s))
}
