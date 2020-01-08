package main

import "fmt"

/*
给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。
例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，
他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个：
{[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}，
{2,3,4,[2,6,2],5,1}，{2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}
*/

// 暴力法
func maxInWindow(s []int, w int) []int {
	var rs []int

	if w == 1 {
		return s
	}

	for k := 0; k < len(s)-w+1; k++ {
		i := 0
		max := s[k]
		for i < w-1 {
			if s[k+i+1] > max {
				max = s[k+i+1]
			}
			i++
		}
		rs = append(rs, max)
	}

	return rs
}

// 使用双向队列保存数组的下标
func maxInWindow2(s []int, w int) []int {
	var rs []int
	var window []int

	if w == 1 {
		return s
	}

	for i := range s {
		// 检查window里存储的下标是否还在当前窗口中有效
		if i >= w && window[0] <= i-w {
			window = window[1:]
		}

		// 如果当前的数字比window里的大 那么就把window小于这个数的数都移除
		for len(window) > 0 && s[window[len(window)-1]] < s[i] {
			window = window[:len(window)-1]
		}

		//把当前数字放入window中
		window = append(window, i)

		//检查是否可以把窗口中的最大值放入结果数组中
		if i >= w-1 {
			rs = append(rs, s[window[0]])
		}
	}

	return rs
}

func main() {
	s := []int{2, 3, 4, 2, 6, 2, 5, 1}
	fmt.Println(maxInWindow(s, 3))
	fmt.Println(maxInWindow2(s, 3))
}
