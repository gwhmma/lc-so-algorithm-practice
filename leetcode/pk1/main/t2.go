package main

import (
	"fmt"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := generateList(add(l1, l2))
	return l
}

func generateList(ns []int) *ListNode {
	l := &ListNode{}
	li := l

	if len(ns) > 0 {
		li.Val = ns[0]
	}
	if len(ns) > 1 {
		for k := 1; k < len(ns); k++ {
			li.Next = &ListNode{Val: ns[k]}
			li = li.Next
		}
	}

	return l
}

func add(l1, l2 *ListNode) []int {
	n1 := getSlice(l1)
	n2 := getSlice(l2)

	var n3 []int
	var le int

	if len(n1) > len(n2) {
		le = len(n1)
	} else {
		le = len(n2)
	}

	i := 0
	k1 := 0
	for {
		if le == 0 && k1 != 0 {
			n3 = append(n3, k1)
			break
		} else if le == 0 && k1 == 0 {
			break
		}
		k := 0
		if i < len(n1) {
			k += n1[i]
		}
		if i < len(n2) {
			k += n2[i]
		}

		k += k1
		tmp := k
		if k >= 10 {
			k %= 10
		}

		n3 = append(n3, k)

		if tmp >= 10 {
			k1 = 1
		} else {
			k1 = 0
		}

		le --
		i ++
	}
	return n3
}

func getSlice(l *ListNode) []int {
	li := l
	var s []int

	for li != nil {
		s = append(s, li.Val)
		li = li.Next
	}
	return s
}

func main() {
	s1 := []int{1}
	s2 := []int{9, 9}

	l1 := generateList(s1)
	l2 := generateList(s2)
	l := addTwoNumbers(l1, l2)

	tmp := l
	for {
		if tmp == nil {
			break
		}
		fmt.Print(" ", tmp.Val)
		tmp = tmp.Next
	}

	// in : [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// in : [5,6,4]
	// correct : [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	// wrong : [2,8,0,4,6,2,5,0,3,0,7,2,4,4,9,6,7,0,5]
}
