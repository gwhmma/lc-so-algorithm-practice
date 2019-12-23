package main

import "fmt"

/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

 

示例 1:

输入：text1 = "abcde", text2 = "ace" 
输出：3  
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func longestCommonSubsequence(t1 string, t2 string) int {
	var dp = make([][]int, len(t1)+1)
	for i := 0; i <= len(t1); i++ {
		dp[i] = make([]int, len(t2)+1)
	}

	for i := 1; i < len(t1)+1; i++ {
		for j := 1; j < len(t2)+1; j++ {
			if t1[i-1:i] == t2[j-1:j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(t1)][len(t2)]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	t1 := "abcd"
	t2 := "a"
	fmt.Println(longestCommonSubsequence(t1, t2))
}
