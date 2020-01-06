package main

import (
	"fmt"
)

/*
请实现一个函数，用来判断一颗二叉树是不是对称的。
注意，如果一个二叉树同此二叉树的镜像是同样的，定义其为对称的
*/

type Tree2 struct {
	val   int
	left  *Tree2
	right *Tree2
}

// 中序遍历 从根节点切分后 2边是逆序对称的
func mirror(t *Tree2) bool {
	var in []int
	key := t.val

	if t == nil || (t.left == nil && t.right == nil) {
		return true
	}

	var inOrder func(*Tree2, *[]int)
	inOrder = func(t *Tree2, in *[]int) {
		if t == nil {
			return
		}

		inOrder(t.left, in)
		*in = append(*in, t.val)
		inOrder(t.right, in)
	}
	inOrder(t, &in)

	idx := 0
	for k, v := range in {
		if v == key {
			idx = k
			break
		}
	}

	left := in[:idx]
	right := in[idx+1:]

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[len(right)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	n1 := &Tree2{val: 1}
	n2 := &Tree2{val: 2}
	n3 := &Tree2{val: 2}
	n4 := &Tree2{val: 3}
	n5 := &Tree2{val: 4}
	n6 := &Tree2{val: 4}
	n7 := &Tree2{val: 3}
	n8 := &Tree2{val: 5}
	n9 := &Tree2{val: 6}
	n10 := &Tree2{val: 7}
	n11 := &Tree2{val: 7}
	n12 := &Tree2{val: 6}
	n13 := &Tree2{val: 5}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	n3.left = n6
	n3.right = n7
	n4.left = n8
	n4.right = n9
	n5.left = n10
	n6.right = n11
	n7.left = n12
	n7.right = n13

	fmt.Println(mirror(n1))
}
