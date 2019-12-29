package main

import "fmt"

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
*/

//首先把二叉树的左右部分分开
//再检查每一部分是否满足二叉搜索树的性质
func VerifySquenceOfBST(s []int) bool {
	if len(s) == 0 {
		return false
	}
	return verify(s, 0, len(s)-1)
}

func verify(s []int, start, end int) bool {
	if start >= end {
		return true
	}

	//首先把二叉树的左右部分分开
	root := s[end]
	i := start
	for {
		if s[i] > root {
			break
		}
		i++
	}

	// 判断右子树是否满足二叉搜索树性质
	j := i
	for j < end {
		if s[j] < root {
			return false
		}
		j++
	}

	return verify(s, start, i-1) && verify(s, i, end-1)
}

func main() {
	s := []int{3,9,5,10,18,15,8}
	fmt.Println(VerifySquenceOfBST(s))
}
