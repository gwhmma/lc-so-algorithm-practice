package main

import "fmt"

/*
汇编语言中有一种移位指令叫做循环左移（ROL），现在有个简单的任务，就是用字符串模拟这个指令的运算结果。
对于一个给定的字符序列S，请你把其循环左移K位后的序列输出。
例如，字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc”。是不是很简单？OK，搞定它！
*/

// 直接切分字符串 然后再拼接
func leftRotate(s string, n int) string {
	s1 := s[:n]
	s = s[n:]
	return s + s1
}

func main() {
	s := "abcXYZdef"
	fmt.Println(leftRotate(s, 3))
}
