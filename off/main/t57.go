package main

import "fmt"

/*
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
例如，链表1->2->3->3->4->4->5 处理后为 1->2->5
*/

type ListNode6 struct {
	val  int
	next *ListNode6
}

func deleteDuplicate1(list *ListNode6) *ListNode6 {
	if list == nil || list.next == nil {
		return list
	}

	if list.val == list.next.val {
		for list.next != nil && list.val == list.next.val {
			list = list.next
		}

		return deleteDuplicate1(list.next)  // 链表1->2->3->3->4->4->5 处理后为 1->2->5
		//return deleteDuplicate1(list)     // 链表1->2->3->3->4->4->5 处理后为 1->2->3->4->5
	} else {
		list.next = deleteDuplicate1(list.next)
	}

	return list
}

func main() {
	node1 := &ListNode6{val: 1}
	node2 := &ListNode6{val: 2}
	node3 := &ListNode6{val: 3}
	node4 := &ListNode6{val: 3 }
	node5 := &ListNode6{val: 4}
	node6 := &ListNode6{val: 4}
	//node7 := &ListNode6{val: 7}
	//node8 := &ListNode6{val: 7}
	//node9 := &ListNode6{val: 7}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	//node6.next = node7
	//node7.next = node8
	//node8.next = node9

	list := deleteDuplicate1(node1)
	for list != nil {
		fmt.Println(list.val)
		list = list.next
	}

}
