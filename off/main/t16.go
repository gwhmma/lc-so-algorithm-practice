package main

import "fmt"

/*
输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
*/

type ListNode3 struct {
	val  int
	next *ListNode3
}

func merge(list1, list2 *ListNode3) *ListNode3 {
	list := &ListNode3{}
	l := list

	if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.val < list2.val {
				tmp := &ListNode3{val:list1.val}
				list.next = tmp
				list1 = list1.next
				list = list.next
			} else {
				tmp := &ListNode3{val:list2.val}
				list.next = tmp
				list2 = list2.next
				list = list.next
			}
		} else if list1 == nil {
			list.next = list2
			break
		} else if list2 == nil {
			list.next = list1
			break
		}
	}

	return l.next
}

func main() {
	l1 := &ListNode3{
		val:  1,
		next: &ListNode3{
			val:  3,
			next: &ListNode3{
				val:  4,
				next: &ListNode3{
					val:  5,
					next: nil,
				},
			},
		},
	}

	l2 := &ListNode3{
		val:  2,
		next: &ListNode3{
			val:  3,
			next: &ListNode3{
				val:  7,
				next: nil,
			},
		},
	}

	list := merge(l1, l2)
	fmt.Println(list)
}
