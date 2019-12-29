package main

import (
	"fmt"
	"math"
)

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。
求该青蛙跳上一个n级的台阶总共有多少种跳法
 */

func jump(n int) int {
	return int(math.Pow(float64(2), float64(n-1)))
}

func main()  {
	fmt.Println(jump(4))
}
