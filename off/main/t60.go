package main

import "fmt"

/*
请实现一个函数按照之字形打印二叉树，即第一行按照从左到右的顺序打印，
第二层按照从右至左的顺序打印，第三行按照从左到右的顺序打印，其他行以此类推


2. 从上到下按层打印二叉树，同一层结点从左至右输出。每一层输出一行
*/

type Tree3 struct {
	val   int
	left  *Tree3
	right *Tree3
}

type queue1 struct {
	val []*Tree3
}

func (q *queue1) push(t *Tree3) {
	q.val = append(q.val, t)
}

func (q *queue1) pop() *Tree3 {
	t := q.val[0]
	q.val = q.val[1:]
	return t
}

func printTree(t *Tree3) {
	var rs [][]int
	var bfs func(t *Tree3)
	q := &queue1{}

	// bfs遍历树 把树的每一层数据放入一个数组中
	q.push(t)
	bfs = func(t *Tree3) {
		for len(q.val) > 0 {
			size := len(q.val)
			var tmp []int
			for i := 0; i < size; i++ {
				node := q.pop()
				tmp = append(tmp, node.val)
				if node.left != nil {
					q.push(node.left)
				}
				if node.right != nil {
					q.push(node.right)
				}
			}

			arr := make([]int, len(tmp))
			copy(arr, tmp)
			rs = append(rs, arr)
		}
	}
	bfs(t)

	// 按规则打印数组
	for k, v := range rs {
		if (k+1)%2 != 0 {
			for _, v := range v {
				fmt.Print(v, " ")
			}
			fmt.Println()
		} else {
			for i := len(v) - 1; i >= 0; i-- {
				fmt.Print(v[i], " ")
			}
			fmt.Println()
		}
	}
}

func main() {
	n1 := &Tree3{val: 1}
	n2 := &Tree3{val: 2}
	n3 := &Tree3{val: 3}
	n4 := &Tree3{val: 4}
	n5 := &Tree3{val: 5}
	n6 := &Tree3{val: 6}
	n7 := &Tree3{val: 7}

	n1.left = n2
	n1.right = n3
	n2.left = n4
	n2.right = n5
	n3.left = n6
	n3.right = n7

	printTree(n1)
}
