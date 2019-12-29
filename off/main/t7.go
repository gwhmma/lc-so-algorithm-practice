package main

import "fmt"

/*
大家都知道斐波那契数列，现在要求输入一个整数n，
请你输出斐波那契数列的第n项（从0开始，第0项为0）。
n<=39
*/

//递归
func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib1(n-1) + fib1(n-2)
}

//非递归

//将每一个n对应的值存储在数组里
func fib2(n int) int {
	var res = []int{1, 1}
	//res = append(res, 1)
	//res = append(res, 1)

	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n-1]
}

//为了节省空间只保存需要的前2个数
func fib3(n int) int {
	b1 := 1
	b2 := 1

	if n < 3 {
		return 1
	}

	for i := 3; i < n; i++ {
		b1, b2 = b2, b1 + b2
	}
	return b1 + b2
}

func main() {
	//fmt.Println(fib1(9))
	//fmt.Println(fib2(9))
	fmt.Println(fib3(10))
}
