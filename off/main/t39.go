package main

import (
	"fmt"
)

/*
输入一棵二叉树，求该树的深度。
从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度
*/

type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

// 树的深度等于 树的节点数 - 叶子节点数 - 1
// dfs 计算出二叉树的节点数和叶子节点数
func treeDepth(t *Tree) int {
	if t == nil {
		return 0
	}
	if t.left == nil && t.right == nil {
		return 1
	}

	var leaf, total int
	var dfs func(t *Tree)

	dfs = func(t *Tree) {
		if t == nil {
			return
		}

		if t.left == nil && t.right == nil {
			leaf++
		}
		total++
		dfs(t.left)
		dfs(t.right)
	}

	dfs(t)

	return total - leaf - 1

}

func main() {
	t1 := &Tree{val: 1}
	t2 := &Tree{val: 2}
	t3 := &Tree{val: 3}
	t4 := &Tree{val: 4}
	t5 := &Tree{val: 5}
	t6 := &Tree{val: 6}
	t7 := &Tree{val: 7}
	t8 := &Tree{val: 8}
	t9 := &Tree{val: 9}
	t10 := &Tree{val: 10}
	t11 := &Tree{val: 11}
	t12 := &Tree{val: 12}
	t13 := &Tree{val: 13}
	t14 := &Tree{val: 13}

	t1.left = t2
	t1.right = t3
	t2.left = t4
	t2.right = t5
	t4.left = t8
	t4.right = t9
	t3.left = t6
	t3.right = t7
	t7.left = t10
	t7.right = t11
	t10.left = t12
	t12.left = t13
	t12.right = t14

	fmt.Println(treeDepth(t1))

}
