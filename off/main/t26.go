package main

import "fmt"

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。
要求不能创建任何新的结点，只能调整树中结点指针的指向。
*/

type TreeNode6 struct {
	val   int
	left  *TreeNode6
	right *TreeNode6
}

type stack4 struct {
	val []*TreeNode6
}

func (s *stack4) push(n *TreeNode6) {
	s.val = append(s.val, n)
}

func (s *stack4) pop() *TreeNode6 {
	n := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return n
}

func (s *stack4) top() *TreeNode6 {
	return s.val[len(s.val)-1]
}

// 二叉搜索树的中序遍历结果就是一个有序的序列
// 所以中序遍历改变节点的指针
// 1 首先将根节点压入栈
// 2 一直压入树的left节点到栈里 知道left节点为空 则开始出栈
// 3 出栈的同时检查这个节点有没有right节点 如果有则重复 2 3

//func convert(tree *TreeNode6) *TreeNode6 {
//	if tree == nil {
//		return tree
//	}
//
//	var s stack4
//	s.push(tree)
//
//	var travel func(s stack4, t *TreeNode6)
//
//	//var var
//
//	travel = func(s stack4, t *TreeNode6) {
//		for len(s.val) > 0 {
//			if s.top().left != nil && s.top().left.right != s.top() {
//				s.push(s.top().left)
//			} else {
//				tmp := s.pop()
//				tmp.right = s.top()
//				s.top().left = tmp
//
//				tmp = s.pop()
//
//				//压入右子树
//				if s.top().right != nil {
//					s.push(s.top().right)
//					travel(s, s.top().right)
//				}
//
//				tmp.right = s.top()
//				s.top().left = tmp
//			}
//		}
//	}
//
//	travel(s, tree)
//	return tree
//}

func convert(tree *TreeNode6) *TreeNode6 {
	var num []*TreeNode6

	if tree == nil {
		return tree
	}

	var travel func(tree *TreeNode6, num *[]*TreeNode6)
	travel = func(tree *TreeNode6, num *[]*TreeNode6) {
		if tree.left != nil {
			travel(tree.left, num)
		}

		*num = append(*num, tree)

		if tree.right != nil {
			travel(tree.right, num)
		}
	}

	travel(tree, &num)

	for i := 1; i < len(num); i++ {
		num[i].left = num[i-1]
		num[i-1].right = num[i]
	}

	return num[0]
}



func main() {
	root := &TreeNode6{val: 10}

	node1 := &TreeNode6{val: 5}
	node2 := &TreeNode6{val: 3}
	node3 := &TreeNode6{val: 1}
	node4 := &TreeNode6{val: 4}
	node5 := &TreeNode6{val: 6}

	node6 := &TreeNode6{val: 15}
	node7 := &TreeNode6{val: 12}
	node8 := &TreeNode6{val: 11}
	node9 := &TreeNode6{val: 13}
	node10 := &TreeNode6{val: 17}
	node11 := &TreeNode6{val: 16}
	node12 := &TreeNode6{val: 18}

	root.left = node1
	root.right = node6
	node1.left = node2
	node1.right = node5
	node2.left = node3
	node2.right = node4
	node6.left = node7
	node6.right = node10
	node7.left = node8
	node7.right = node9
	node10.left = node11
	node11.right = node12

	t := convert(root)
	fmt.Println(t)

}
