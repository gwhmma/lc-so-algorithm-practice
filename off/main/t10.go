package main

import "fmt"

/*
我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。
请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
*/

func RectCover(n int) int {
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
	fmt.Println(RectCover(4))
}
