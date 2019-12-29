package main

import "fmt"

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，
则重建二叉树并返回
*/

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// pre前序遍历   in中序遍历
// 1.根据前序序列的第一个元素得到二叉树的根节点
// 2.根据1中的元素把中序序列分为左右2个子序列部分
// 3.对左右部分的元素重复进行1 2 步骤
func reBuild(pre, in []int) *TreeNode {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}

	tree := &TreeNode{}
	tree.val = pre[0]
	for i := 0; i < len(in); i++ {
		if in[i] == pre[0] {
			tree.left = reBuild(pre[1:i+1], in[:i])
			tree.right = reBuild(pre[i+1:], in[i+1:])
			break
		}
	}
	return tree
}

func main() {
	pre := []int{1, 3, 4, 5, 6, 11, 10}
	in := []int{4, 3, 5, 1, 11, 6, 10}
	tree := reBuild(pre, in)
	fmt.Println(tree)
}
