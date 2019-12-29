package main

import "fmt"

/*
从上往下打印出二叉树的每个节点，同层节点从左至右打印。
*/

type queue struct {
	val []*TreeNode4
}

func (q *queue) push(n *TreeNode4) {
	q.val = append(q.val, n)
}

func (q *queue) pop() *TreeNode4 {
	t := q.val[0]
	q.val = q.val[1:]
	return t
}

func (q *queue) top() *TreeNode4 {
	return q.val[0]
}

type TreeNode4 struct {
	val   int
	left  *TreeNode4
	right *TreeNode4
}

// 1. 首先将根节点放入队列中
// 2. 当队列不为空时 取出队列的第一个元素 同时检查该元素是否还有子节点
// 3. 如果该节点还有子节点则将其子节点放入队列中
// 4. 2 3 步骤 直到队列为空

func print(tree *TreeNode4) []int {
	var res []int
	q := queue{}

	if tree == nil {
		return res
	}

	q.push(tree)
	for len(q.val) > 0 {
		tmp := q.pop()
		res = append(res, tmp.val)
		if tmp.left != nil {
			q.push(tmp.left)
		}
		if tmp.right != nil {
			q.push(tmp.right)
		}
	}

	return res
}

func main() {
	tree := &TreeNode4{
		val: 1,
		left: &TreeNode4{
			val: 2,
			left: &TreeNode4{
				val:   4,
				left:  nil,
				right: nil,
			},
			right: &TreeNode4{
				val:   5,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeNode4{
			val: 3,
			left: &TreeNode4{
				val:   6,
				left:  nil,
				right: nil,
			},
			right: &TreeNode4{
				val:   7,
				left:  nil,
				right: nil,
			},
		},
	}

	fmt.Println(print(tree))

}
