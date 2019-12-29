package main

import "fmt"

/*
操作给定的二叉树，将其变换为源二叉树的镜像。

二叉树的镜像定义：源二叉树
    	    8
    	   /  \
    	  6   10
    	 / \  / \
    	5  7 9 11
    	镜像二叉树
    	    8
    	   /  \
    	  10   6
    	 / \  / \
    	11 9 7  5
*/

type TreeNode3 struct {
	val   int
	left  *TreeNode3
	right *TreeNode3
}

//递归对左右子树进行旋转
func Mirror(t *TreeNode3)  {
	if t == nil {
		return
	}

	tmp := t.left
	t.left = t.right
	t.right = tmp
	Mirror(t.left)
	Mirror(t.right)
}

func main() {
	tree := &TreeNode3{
		val:   8,
		left:  &TreeNode3{
			val:   6,
			left:  &TreeNode3{
				val:   5,
				left:  nil,
				right: nil,
			},
			right: &TreeNode3{
				val:   7,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeNode3{
			val:   10,
			left:  &TreeNode3{
				val:   9,
				left:  nil,
				right: nil,
			},
			right: &TreeNode3{
				val:   11,
				left:  nil,
				right: &TreeNode3{
					val:   4,
					left:  nil,
					right: nil,
				},
			},
		},
	}

	Mirror(tree)
	fmt.Println(tree)

}
