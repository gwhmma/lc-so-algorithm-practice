package main

import "fmt"

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级。
求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）
*/

func JumpFloor(n int) int {
	if n < 3 {
		return n
	}

	b1 := 1
	b2 := 2

	for i := 3; i < n; i++ {
		b1, b2 = b2, b1+b2
	}
	return b1 + b2
}

func main() {
	fmt.Println(JumpFloor(5))
}
