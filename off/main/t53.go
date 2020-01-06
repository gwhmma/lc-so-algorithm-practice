package main

import "fmt"

/*
请实现一个函数用来匹配包括'.'和'*'的正则表达式。
模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（包含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。
例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab*a"均不匹配
*/

func match(s, pattern []string) bool {
	if len(s) == 0 || len(pattern) == 0 {
		return false
	}

	return matchStr(s, 0, pattern, 0)
}

func matchStr(str []string, strIndex int, pattern []string, patIndex int) bool {
	// str 或者 pattern 到末尾 匹配成功
	if strIndex == len(str) && patIndex == len(pattern) {
		return true
	}
	// pattern先到达末尾 匹配失败
	if strIndex != len(str) && patIndex == len(pattern) {
		return false
	}

	//模式pattern第二个字符是 * ，且字符串第一个和模式第一个匹配 则分为3种匹配模式  如果不匹配模式后移2位
	if patIndex+1 < len(pattern) && pattern[patIndex+1] == "*" {
		if (strIndex != len(str) && pattern[patIndex] == str[strIndex]) || (pattern[patIndex] == "." && strIndex != len(str)) {
			return matchStr(str, strIndex, pattern, patIndex+2) || // 模式后移2位 x* 匹配0个字符
				matchStr(str, strIndex+1, pattern, patIndex+2) || // 模式匹配1个字符
				matchStr(str, strIndex+1, pattern, patIndex) // *匹配一个 再匹配str中的下一个
		}
	} else {
		return matchStr(str, strIndex, pattern, patIndex+2)
	}

	// 模式第二个不是* 且字符串第一个和模式第一个匹配  则都后移1位  否则直接返回false
	if (strIndex != len(str) && pattern[patIndex] == str[strIndex]) || (pattern[patIndex] == "." && strIndex != len(str)) {
		return matchStr(str, strIndex, pattern, patIndex)
	}
	return false
}

func main() {
	str := []string{"a", "a", "a"}
	pattern := []string{"a", ".", "a"}
	fmt.Println(match(str, pattern))
}
