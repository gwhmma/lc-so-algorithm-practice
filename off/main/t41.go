package main

import "fmt"

/*
一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字
*/

// 将数字放入map中 k为该数字 v为这个数字出现的次数
func notRepeat(a []int) []int {
	var m = make(map[int]int)
	var rs []int

	for _, v := range a {
		m[v]++
	}

	for k, v := range m {
		if len(rs) == 2 {
			break
		}

		if v == 1 {
			rs = append(rs, k)
		}
	}

	return rs
}

func main() {
	a := []int{1,2,2,1,3,4,5,4,6,6,7,8,9,9,8,7}
	fmt.Println(notRepeat(a))
}
