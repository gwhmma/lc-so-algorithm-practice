package main

import "fmt"

/*
输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
*/

func NumberOf1(n int) int {
	count := 0
	first := false

	if n > 0 {
		for n > 0 {
			if n&1 == 1 {
				count++
			}
			n >>= 1
		}
	} else {
		n = -n
		for n > 0 {
			if n&1 == 0 {
				count++
			} else if n&1 == 1 && !first {
				count++
			}
			first = true
			n >>= 1
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(NumberOf1(15))
	fmt.Println(NumberOf1(-15))
}
