package main

import "fmt"

/*
在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,
并返回它的位置, 如果没有则返回 -1（需要区分大小写）
*/

//将首先遍历字符串 将每个字符出现的次数进行统计写入map
//再次遍历字符串 以每个字符去map中匹配第一个次数为1的字符
func firstNoRepeat(s string) int {
	if len(s) <= 1 {
		return len(s) - 1
	}

	var m = make(map[byte]int)
	for _, v := range []byte(s) {
		m[v]++
	}

	for k, v := range []byte(s) {
		if m[v] == 1 {
			fmt.Println(string(v))
			return k
		}
	}

	return -1
}

// map是无序的
// 所以这题参见 t55的代码

func main() {
	s := "abcbcbcxchjbkjhkjukp"
	fmt.Println(firstNoRepeat(s))
}
