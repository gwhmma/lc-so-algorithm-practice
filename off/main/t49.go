package main

import (
	"fmt"
	"strconv"
)

/*
写一个函数，求两个整数之和，要求在函数体内不得使用+、-、*、/四则运算符号
*/

func sum2(a, b int) int64 {
	rs := ""
	f := false

	for a > 0 || b > 0 {
		if (a&1) == 1 && (b&1) == 1 {
			if f {
				rs = fmt.Sprintf("%d%s", 1, rs)
				f = false
			} else {
				rs = fmt.Sprintf("%d%s", 0, rs)
			}
			f = true
			a >>= 1
			b >>= 1
		} else if (a&1) == 1 || (b&1) == 1 {
			if f {
				rs = fmt.Sprintf("%d%s", 0, rs)
			} else {
				rs = fmt.Sprintf("%d%s", 1, rs)
			}
			a >>= 1
			b >>= 1
		} else {
			if f {
				rs = fmt.Sprintf("%d%s", 1, rs)
				f = false
			} else {
				rs = fmt.Sprintf("%d%s", 0, rs)
			}
			a >>= 1
			b >>= 1
		}
	}

	if f {
		rs = fmt.Sprintf("%d%s", 1, rs)
	}

	s, _ := strconv.ParseInt(rs, 2, 64)
	return s
}

func main() {
	fmt.Println(sum2(5, 9))
}
