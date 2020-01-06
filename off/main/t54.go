package main

import "fmt"

/*
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，字符串"+100","5e2","-123","3.1416"和"-1E-16"都表示数值。
但是"12e","1a3.14","1.2.3","+-5"和"12e+4.3"都不是
*/

func checkNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	bt := []byte(s)
	for k, v := range bt {
		switch v {
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57: // 0 - 9
		case 43: // +
			// + 只能出现在第一位
			if k > 0 {
				return false
			}
		case 45: // -
			// - 可以出现在第一位  或者 - 之前只能是 E e
			if k > 0 && (bt[k-1] != 69 && bt[k-1] != 101) {
				return false
			}
		case 46: // .
			// .不能出现在第一位   .只能出现在数字之间
			if k == 0 || k == len(bt)-1 || (k > 0 && (bt[k-1] < 48 && bt[k-1] > 57)) {
				return false
			}
		case 69, 101: // E  e
			// E e 不能出现在第一位   E e 前面只能是数字  E e 后面可以是数字和-
			if k == 0 || (k > 0 && (bt[k-1] < 48 || bt[k-1] > 57)) {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func main() {
	s := "-1.2E-3.1"
	fmt.Println(checkNumber(s))
}
