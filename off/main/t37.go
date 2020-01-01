package main

import "fmt"

/*
输入两个链表，找出它们的第一个公共结点
*/

type ListNode4 struct {
	val  int
	next *ListNode4
}

// 首先计算出2个链表的长度
// 较长的链表先走 2个链表的 "长度差" 步
// 然后2个链表在同时往后走 如果他们的节点相同 则这个节点为第一个公共节点

func firstCommonNode(l1, l2 *ListNode4) *ListNode4 {
	len1 := 0
	len2 := 0
	t1 := l1
	t2 := l2

	for t1 != nil {
		len1++
		t1 = t1.next
	}
	for t2 != nil {
		len2++
		t2 = t2.next
	}

	if len1 > len2 {
		abs := len1 - len2
		node := div(l1, l2, abs)
		return node
	} else {
		abs := len2 - len1
		node := div(l2, l1, abs)
		return node
	}
}

func div(l1, l2 *ListNode4, abs int) *ListNode4 {
	for abs > 0 {
		l1 = l1.next
		abs--
	}

	for l1 != nil && l2 != nil {
		if l1 == l2 {
			return l1
		}
		l1 = l1.next
		l2 = l2.next
	}
	return nil
}

func main() {
	node1 := &ListNode4{val: 1}
	node2 := &ListNode4{val: 2}
	node3 := &ListNode4{val: 3}
	node4 := &ListNode4{val: 4}
	node5 := &ListNode4{val: 5}
	node6 := &ListNode4{val: 6}

	node7 := &ListNode4{val: 11}
	node8 := &ListNode4{val: 12}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6

	node7.next = node8
	node8.next = node5

	fmt.Println(firstCommonNode(node1, node7))
}
