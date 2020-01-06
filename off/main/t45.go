package main

import (
	"fmt"
	"strings"
)

/*
牛客最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。
同事Cat对Fish写的内容颇感兴趣，有一天他向Fish借来翻看，但却读不懂它的意思。
例如，“student. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，
正确的句子应该是“I am a student.”。Cat对一一的翻转这些单词顺序可不在行，你能帮助他么？
*/

// 将字符串通过 空格 " " split进一个数组 然后逆序遍历这个数组取出每个单词
func reverseSentence(s string) string {
	var rs string

	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		rs += words[i] + " "
	}

	return strings.TrimRight(rs, " ")
}

func main() {
	s := "apple. the wants he and orange, eat to like I'd"
	fmt.Println(reverseSentence(s))
}
