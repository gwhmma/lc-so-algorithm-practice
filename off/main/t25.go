package main

/*
输入一个复杂链表（每个节点中有节点值，以及两个指针，一个指向下一个节点，
另一个特殊指针指向任意一个节点），返回结果为复制后复杂链表的head。
（注意，输出结果中请不要返回参数中的节点引用，否则判题程序会直接返回空）
*/

type RandomListNode struct {
	val    int
	next   *RandomListNode
	random *RandomListNode
}

//1 将链表每个节点复制一个跟在原节点之后
//2 新节点的随机指针就在原节点随机指针的下一个节点
//3 拆分新节点 奇数个是原节点 偶数个是新节点
func clone(list *RandomListNode) *RandomListNode {
	if list == nil || list.next == nil {
		return list
	}

	//1 将链表每个节点复制一个跟在原节点之后
	tmp := list
	tmp1 := tmp
	for tmp != nil {
		t := tmp.next
		//注意 这里新加的一个节点需要重新new一个
		node := &RandomListNode{val: tmp.val}
		tmp.next = node
		node.next = t
		tmp = tmp.next.next
	}

	//2 新节点的随机指针就在原节点随机指针的下一个节点
	tmp2 := tmp1
	for tmp1 != nil {
		if tmp1.random != nil {
			tmp1.next.random = tmp1.random.next
		} else {
			tmp1.next.random = nil
		}
		tmp1 = tmp1.next.next
	}

	//3 拆分新节点 奇数个是原节点 偶数个是新节点
	tmp3 := tmp2.next
	res := tmp3
	for tmp3.next != nil {
		tmp3.next = tmp3.next.next
		tmp3 = tmp3.next
	}

	return res
}

func main() {
	node1 := &RandomListNode{
		val:    1,
		next:   nil,
		random: nil,
	}

	node2 := &RandomListNode{
		val:    2,
		next:   nil,
		random: nil,
	}

	node3 := &RandomListNode{
		val:    3,
		next:   nil,
		random: nil,
	}

	node1.next = node2
	node2.next = node3
	node1.random = node3
	node3.random = node1

	clone(node1)
}
