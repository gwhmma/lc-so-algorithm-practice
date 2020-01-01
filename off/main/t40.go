package main

import "fmt"

/*
输入一棵二叉树，判断该二叉树是否是平衡二叉树
*/

type TreeNode7 struct {
	val   int
	left  *TreeNode7
	right *TreeNode7
}

// 检查每二叉树的每条路径深度  深度差值 < 2 则是平衡二叉树
func checkAVL(tree *TreeNode7) bool {
	var blance func(t *TreeNode7) (int, bool)
	blance = func(t *TreeNode7) (i int, b bool) {
		if t == nil {
			return 0, true
		}

		ld, lb := blance(t.left)
		rd, rb := blance(t.right)
		return max(ld, rd) + 1, lb && rb && abs(ld-rd) < 2
	}

	_, rs := blance(tree)
	return rs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	n1 := &TreeNode7{val: 1}
	n2 := &TreeNode7{val: 2}
	n3 := &TreeNode7{val: 3}
	n4 := &TreeNode7{val: 4}
	n5 := &TreeNode7{val: 5}
	//n6 := &TreeNode7{val: 6}
	//n7 := &TreeNode7{val: 7}
	//n8 := &TreeNode7{val: 8}
	//n9 := &TreeNode7{val: 9}
	//n10 := &TreeNode7{val: 10}
	//n11 := &TreeNode7{val: 11}
	//n12 := &TreeNode7{val: 12}
	//n13 := &TreeNode7{val: 13}
	//n14 := &TreeNode7{val: 14}
	//n15 := &TreeNode7{val: 15}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	//n4.left = n8
	//n4.right = n9
	//n8.right = n15

	//n3.left = n6
	//n4.left = n8
	//n8.right = n15
	//n3.right = n7
	//n7.left = n10
	//n7.right = n11
	//n10.left = n12
	//n12.left = n13
	//n12.right = n14

	fmt.Println(checkAVL(n1))

}
