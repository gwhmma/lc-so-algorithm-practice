package main

import "fmt"

/*
给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null
*/

type ListNode5 struct {
	val  int
	next *ListNode5
}

// 快慢指针  慢指针一次走一步 快指针一次走2不
// 当2个指针第一次相遇时 让慢指针继续走 同时从链表头开始走 2个指针相遇的节点就是入口节点
// 数学推理😎
func entryOfLoop(list *ListNode5) *ListNode5 {
	slow := list
	fast := list
	l := list

	if list == nil {
		return nil
	}

	for fast != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			for {
				if l == slow {
					return l
				}
				l = l.next
				slow = slow.next
			}
		}
	}

	return nil
}

func main() {
	node1 := &ListNode5{val: 1}
	node2 := &ListNode5{val: 2}
	node3 := &ListNode5{val: 3}
	node4 := &ListNode5{val: 4}
	node5 := &ListNode5{val: 5}
	node6 := &ListNode5{val: 6}
	node7 := &ListNode5{val: 7}
	node8 := &ListNode5{val: 8}
	node9 := &ListNode5{val: 9}
	node10 := &ListNode5{val: 10}
	node11 := &ListNode5{val: 11}
	node12 := &ListNode5{val: 12}
	node13 := &ListNode5{val: 13}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7
	node7.next = node8
	node8.next = node9
	node9.next = node10
	node10.next = node11
	node11.next = node12
	node12.next = node13
	node13.next = node9

	fmt.Println(entryOfLoop(node1).val)
}
