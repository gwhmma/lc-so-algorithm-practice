package main

import "fmt"

/*
把只包含质因子2、3和5的数称作丑数（Ugly Number）
例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第N个丑数
*/

func uglyNumber(n int) int {
	var ugly []int

	if n < 7 {
		return n
	}

	ugly = append(ugly, 1)
	p2 := 0
	p3 := 0
	p5 := 0
	uglyNum := 0

	for len(ugly) < n {
		uglyNum = min(2*ugly[p2], min(3*ugly[p3], 5*ugly[p5]))
		if uglyNum == ugly[p2]*2 {
			p2++
		}
		if uglyNum == ugly[p3]*3 {
			p3++
		}
		if uglyNum == ugly[p5]*5 {
			p5++
		}
		ugly = append(ugly, uglyNum)
	}
	return ugly[len(ugly)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	fmt.Println(uglyNumber(12000))
}
