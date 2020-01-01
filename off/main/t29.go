package main

import (
	"fmt"
	"sort"
)

/*
输入n个整数，找出其中最小的K个数。
例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4
*/

func minK(a []int, k int) []int {
	sort.Ints(a)
	return a[:k]
}

func minK2(a []int, k int) []int {
	var res []int
	for _, v := range a {
		if len(res) < k {
			res = append(res, v)
		} else {
			check(&res, v, k)
		}
	}
	return res
}

// TODO 最大/小堆 优先队列 将k个数放入 之后就不用排序  可以直接替换最大的数
func check(res *[]int, n, k int) {
	sort.Ints(*res)
	if n < (*res)[k-1] {
		(*res)[k-1] = n
	}
}

func main() {
	a := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(minK(a, 5))
	fmt.Println(minK2(a, 5))
}
