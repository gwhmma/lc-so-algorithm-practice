package main

import "fmt"

/*
输入一个链表，按链表从尾到头的顺序返回一个ArrayList
*/

type ListNode struct {
	val  int
	next *ListNode
}

func printListFromTailToHead(list *ListNode) []int {
	var res []int
	head := list

	for head != nil {
		res = append(res, head.val)
		head = head.next
	}

	return res
}

func main() {
	list := &ListNode{
		val:  0,
		next: &ListNode{
			val:  1,
			next: &ListNode{
				val:  2,
				next: &ListNode{
					val:  3,
					next: nil,
				},
			},
		},
	}

	fmt.Println(printListFromTailToHead(list))
}
