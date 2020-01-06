package main

import "fmt"

/*
在一个长度为n的数组里的所有数字都在0到n-1的范围内。
数组中某些数字是重复的，但不知道有几个数字是重复的。
也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是第一个重复的数字2
*/


//用一个map存储每个数字出现的次数
//存储的同事检查这个数字出现的次数 如果大于1就直接返回这个数字
func duplicate(a []int) int {
	var dup = make(map[int]int)

	for _, v := range a {
		dup[v]++
		if dup[v] > 1 {
			return v
		}
	}

	return 0
}

func main() {
	a := []int{2, 3, 1, 0, 3, 2, 5, 3}
	fmt.Println(duplicate(a))
}
