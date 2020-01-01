package main

import (
	"fmt"
	"strconv"
)

/*
输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323
*/

//1 首先从数组第一位开始  将第一位于第二位数字进行拼接
//2 选取最小的那一个 例如 (3, 32) ---> (332, 323) 则选取323
//3 继续使用这个最小的拼接数字和后续数字进行拼接 重复1 2 步骤
func minNum(s []int) {
	rs := strconv.Itoa(s[0])
	for i := 1; i < len(s); i++ {
		s1 := rs + strconv.Itoa(s[i])
		s2 := strconv.Itoa(s[i]) + rs

		a1, _ := strconv.Atoi(s1)
		a2, _ := strconv.Atoi(s2)
		if a1 > a2 {
			rs = strconv.Itoa(a2)
		} else {
			rs = strconv.Itoa(a1)
		}
	}
	fmt.Println(rs)
}

func main() {
	s := []int{3, 32, 321,321}
	minNum(s)
}
