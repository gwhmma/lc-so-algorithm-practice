package main

import "fmt"

/*
请实现一个函数用来找出字符流中第一个只出现一次的字符。
例如，当从字符流中只读出前两个字符"go"时，第一个只出现一次的字符是"g"。
当从该字符流中读出前六个字符“google"时，第一个只出现一次的字符是"l"。
输出描述:
如果当前字符流没有存在出现一次的字符，返回#字符。
*/

// 将每个字符出现的次数记录在map里
// 因为map无序 所以需要用原来的字符串为顺序来从map里取第一个无重复的字符
func firstNotRepeat(s string) string {
	var m = make(map[byte]int)

	bt := []byte(s)
	for _, v := range bt {
		m[v]++
	}

	for _, v := range bt {
		if m[v] == 1 {
			return string(v)
		}
	}

	return "#"
}

func main() {
	s := "google"
	fmt.Println(firstNotRepeat(s))
}
