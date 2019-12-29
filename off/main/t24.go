package main

import "fmt"

/*
输入一颗二叉树的跟节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
(注意: 在返回值的list中，数组长度大的数组靠前)
*/

type List struct {
	list []int
}

type stack3 struct {
	val []*TreeNode5
}

func (s *stack3) push(t *TreeNode5) {
	s.val = append(s.val, t)
}

func (s *stack3) pop() *TreeNode5 {
	t := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return t
}

type TreeNode5 struct {
	val   int
	left  *TreeNode5
	right *TreeNode5
}

//func FindPath(tree *TreeNode5, n int) []List {
//	var list []List
//	s := stack3{}
//
//	if tree.val == n {
//		tmp := []int{n}
//		l := List{tmp}
//		list = append(list, l)
//		return list
//	}
//	if tree.val > n {
//		return list
//	}
//
//	s.push(tree)
//	add := 0
//	l := List{}
//	for len(s.val) > 0 {
//		tmp := s.pop()
//		add += tmp.val
//		if add == n {
//			l.list = append(l.list, tmp.val)
//			list = append(list, l)
//			l.list = l.list[:len(l.list)-1]
//			add -= tmp.val
//		} else if add > n {
//			l.list = l.list[:len(l.list)-1]
//			add -= tmp.val
//		} else {
//			l.list = append(l.list, tmp.val)
//			if tmp.right != nil {
//				s.push(tmp.right)
//			}
//			if tmp.left != nil {
//				s.push(tmp.left)
//			}
//		}
//
//	}
//
//}

// 递归
func FindPath(tree *TreeNode5, n int) [][]int {
	var res [][]int
	var list []int

	var dfs func(tree *TreeNode5, sum int)
	dfs = func(tree *TreeNode5, sum int) {
		if tree == nil {
			return
		}
		rest := sum - tree.val

		if rest == 0 && tree.left == nil && tree.right == nil {
			list = append(list, tree.val)
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			list = list[:len(list)-1]
		}

		list = append(list, tree.val)
		dfs(tree.left, rest)
		dfs(tree.right, rest)
		list = list[:len(list)-1]
	}

	dfs(tree, n)
	return res
}

func main() {
	t := &TreeNode5{
		val:   1,
		left:  &TreeNode5{
			val:   2,
			left:  &TreeNode5{
				val:   5,
				left:  nil,
				right: nil,
			},
			right: &TreeNode5{
				val:   5,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeNode5{
			val:   3,
			left:  &TreeNode5{
				val:   4,
				left:  nil,
				right: nil,
			},
			right: &TreeNode5{
				val:   5,
				left:  nil,
				right: nil,
			},
		},
	}

	fmt.Println(FindPath(t, 8))
}
