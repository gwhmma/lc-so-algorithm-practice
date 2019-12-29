package main

import "fmt"

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，
所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。
*/

//把奇数放一个数组 偶数放一个数组 最后合并
func reOrderArray1(s []int) {
	var s1 []int
	var s2 []int

	for _, v := range s {
		if v%2 == 0 {
			s1 = append(s1, v)
		} else {
			s2 = append(s2, v)
		}
	}
	res := append(s2,s1...)
	fmt.Println(res)
}

//
func reOrderArray2(s []int) {

}

func main() {
	s := []int{1,4,3,6,5,7,4,2}
	reOrderArray1(s)
}
