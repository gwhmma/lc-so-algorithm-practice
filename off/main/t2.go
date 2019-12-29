package main

import (
	"fmt"
	"strings"
)

/*
请实现一个函数，将一个字符串中的每个空格替换成“%20”。
例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy
 */

func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ReplaceAll(s, " ", "%")
}

func main()  {
	s := "I  love you "
	fmt.Println(replaceSpace(s))
}
