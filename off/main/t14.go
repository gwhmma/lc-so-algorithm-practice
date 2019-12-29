package main

import "fmt"

/*
输入一个链表，输出该链表中倒数第k个结点
*/

type ListNode1 struct {
	val  int
	next *ListNode1
}

// 首先得到list的长度  则第倒数第k个节点就为 第 len - k + 1 个节点
func FindKthToTail1(list *ListNode1, k int) *ListNode1 {
	length := 0
	cur := list
	//var node *ListNode1

	for cur != nil {
		length++
		cur = cur.next
	}

	t := length - k + 1
	for t > 0 {
		if t == 1 {
			return &ListNode1{val: list.val}
		} else {
			list = list.next
			t--
		}
	}
	return nil
}

// 快慢指针
// 1 先让第一个指针走k+1步
// 2 第二个指针从头开始走 第一个指针继续走
// 3 当第一个指针到达最后一个节点时  第二个指针的next即为我们需要的
func FindKthToTail2(list *ListNode1, k int) *ListNode1 {
	p1 := list
	p2 := list

	k1 := k + 1
	for k1 > 0 {
		p1 = p1.next
		k1--
	}

	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}

	return &ListNode1{val:p2.next.val}
}

func main() {
	list := &ListNode1{val: 1,
		next: &ListNode1{val: 2,
			next: &ListNode1{
				val: 3,
				next: &ListNode1{
					val: 4,
					next: &ListNode1{
						val: 5,
						next: &ListNode1{
							val: 6,
							next: &ListNode1{
								val:  7,
								next: nil,
							},
						},
					},
				},
			}}}
	node := FindKthToTail1(list, 3)
	node2 := FindKthToTail2(list, 3)
	fmt.Println(node)
	fmt.Println(node2)
}
