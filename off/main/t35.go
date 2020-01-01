package main

import "fmt"

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。
即输出P%1000000007

题目保证输入的数组中没有的相同的数字

数据范围：

	对于%50的数据,size<=10^4

	对于%75的数据,size<=10^5

	对于%100的数据,size<=2*10^5

示例1
输入
1,2,3,4,5,6,7,0
输出
7
*/

//归并排序
//排序的过程中统计逆序对
func inversePairs(s []int) int {
	rs := 0
	merge1(s, &rs)
	return rs
}

func merge1(s []int, count *int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]
	return sort(merge1(left, count), merge1(right, count), count)
}

func sort(left, right []int, count *int) []int {
	var tmp []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			tmp = append(tmp, left[0])
			left = left[1:]
		} else {
			tmp = append(tmp, right[0])
			right = right[1:]
			*count++
		}
	}

	if len(left) > 0 {
		tmp = append(tmp, left...)
	}
	if len(right) > 0 {
		tmp = append(tmp, right...)
	}
	return tmp
}

func main() {
	s := []int{1,2,3,4,5,6,7,0}
	fmt.Println(inversePairs(s))
}
