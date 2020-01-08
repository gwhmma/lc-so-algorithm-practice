package main

import "fmt"

/*
给定一棵二叉搜索树，请找出其中的第k小的结点。
例如， （5，3，7，2，4，6，8）    中，按结点数值大小顺序第三小结点的值为4
*/

type Tree5 struct {
	val   int
	left  *Tree5
	right *Tree5
}

// 二叉搜索树最小的值是中序遍历的第一个节点(最左下的叶子节点 如果有的话)
func minKtn(tree *Tree5, k int) *Tree5 {
	if tree == nil {
		return nil
	}

	var num []*Tree5
	var inOrder func(tree *Tree5)

	inOrder = func(tree *Tree5) {
		if tree == nil {
			return
		}

		//中序遍历二叉搜索树 并将节点放入slice中 当slice size = k 这是 slice最后一个元素就是第k小的节点
		if len(num) < k {
			inOrder(tree.left)
			num = append(num, tree)
			inOrder(tree.right)
		} else {
			return
		}
	}
	inOrder(tree)

	return num[k-1]
}

func main() {
	n1 := &Tree5{val: 5}
	n2 := &Tree5{val: 3}
	n3 := &Tree5{val: 7}
	n4 := &Tree5{val: 2}
	n5 := &Tree5{val: 4}
	n6 := &Tree5{val: 6}
	n7 := &Tree5{val: 8}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	n3.left = n6
	n3.right = n7

	node := minKtn(n1, 5)
	fmt.Println(node.val)
}
