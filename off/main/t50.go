package main

import "fmt"

/*
将一个字符串转换成一个整数，要求不能使用字符串转换整数的库函数。 数值为0或者字符串不是一个合法的数值则返回0
输入描述:
输入一个字符串,包括数字字母符号,可以为空
输出描述:
如果是合法的数值表达则返回该数字，否则返回0

示例1
输入
+2147483647
    1a33

输出
2147483647
    0
*/

// asc II码    0 : 48    9 : 57   + : 43  - : 45
func trans(s string) int {
	rs := 0
	negtive := false

	if s[:1] == "-" {
		negtive = true
		s = s[1:]
	} else if s[:1] == "+" {
		s = s[1:]
	}

	byt := []byte(s)
	j := 1
	for i := len(byt) - 1; i >= 0; i-- {
		if byt[i] >= 48 && byt[i] <= 57 {
			switch byt[i] {
			case 48:
			case 49:
				rs += j
			case 50:
				rs += 2 * j
			case 51:
				rs += 3 * j
			case 52:
				rs += 4 * j
			case 53:
				rs += 5 * j
			case 54:
				rs += 6 * j
			case 55:
				rs += 7 * j
			case 56:
				rs += 8 * j
			case 57:
				rs += 9 * j
			}
			j *= 10
		} else {
			return 0
		}
	}

	if negtive {
		rs = -rs
	}
	return rs
}

func main() {
	fmt.Println(trans("+2147483647"))
}
