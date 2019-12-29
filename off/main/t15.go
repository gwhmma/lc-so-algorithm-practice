package main

import "fmt"

/*
输入一个链表，反转链表后，输出新链表的表头。
*/

type ListNode2 struct {
	val  int
	next *ListNode2
}

func reverse(list *ListNode2) *ListNode2 {
	if list == nil || list.next == nil {
		return list
	}

	pre := list
	cur := pre.next
	tmp := cur.next
	pre.next = nil

	for cur != nil {
		cur.next = pre
		pre = cur
		cur = tmp
		if tmp != nil {
			tmp = tmp.next
		}
	}
	return pre
}

func main() {
	list := &ListNode2{
		val: 1,
		next: &ListNode2{
			val: 2,
			next: &ListNode2{
				val: 3,
				next: &ListNode2{
					val:  4,
					next: nil,
				},
			},
		},
	}

	r := reverse(list)
	fmt.Println(r)
}
