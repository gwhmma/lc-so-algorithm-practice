package main

import "fmt"

/*
给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回。
注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针
*/

type Tree1 struct {
	val    int
	left   *Tree1
	right  *Tree1
	parent *Tree1
}

// 中序遍历 : 左根右
func getNext(t, node *Tree1) *Tree1 {
	if t == nil || node == nil {
		return nil
	}

	// 判断当前给的节点是左节点 右节点  还是根节点

	if node.left == nil && node.right == nil {
		// 左叶子节点
		if node.parent.left == node {
			return node.parent
		} else if node.parent.right == node && node.parent.parent.left == node.parent {
			// 右叶子节点
			return node.parent.parent
		}
	} else {
		if node.right != nil {
			return node.right
		} else if node.parent != nil && node.parent.right != node {
			return node.parent
		}
	}
	return nil
}

func main() {
	node1 := &Tree1{val: 1}
	node2 := &Tree1{val: 2}
	node3 := &Tree1{val: 3}
	node4 := &Tree1{val: 4}
	node5 := &Tree1{val: 5}
	node6 := &Tree1{val: 6}
	node7 := &Tree1{val: 7}
	node8 := &Tree1{val: 8}
	node9 := &Tree1{val: 9}

	node1.left = node2
	node1.right = node3
	node2.parent = node1
	node3.parent = node1
	node2.left = node4
	node2.right = node5
	node4.parent = node2
	node5.parent = node2
	node3.left = node6
	node3.right = node7
	node6.parent = node3
	node7.parent = node3
	node4.left = node8
	node4.right = node9
	node8.parent = node4
	node9.parent = node4

	fmt.Println(getNext(node1, node9))
}
