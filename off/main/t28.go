package main

import (
	"fmt"
	"sort"
)

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如输入一个长度为9的数组{1,2,3,2,2,2,5,4,2}。
由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。
*/

func more(s []int) int {
	sort.Ints(s)
	key := s[len(s)/2]
	i := 0
	for _, v := range s {
		if v == key {
			i++
		}
	}

	if i > len(s)/2 {
		return key
	}
	return 0
}

func main() {
	s := []int{1, 2, 3, 2, 2, 2, 5, 4,2}
	fmt.Println(more(s))
}
