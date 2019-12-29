package main

import "fmt"

/*
用两个栈来实现一个队列，完成队列的Push和Pop操作。
队列中的元素为int类型
*/

type stack struct {
	val []int
}

func (s *stack) push(n int) {
	s.val = append(s.val, n)
}

func (s *stack) pop() int {
	n := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return n
}

// 1. push操作直接将数压入s1
// 2. pop操作时首先判断s2是否为空，如果不为空直接从s2中pop数据 如果为空则现将s1中数据push到s2中在进行pop

func push(val int) {
	s1.push(val)
}

func pop() int {
	n := 0

	if len(s2.val) == 0 {
		for len(s1.val) != 0 {
			s2.push(s1.pop())
		}
		n = s2.pop()
	} else {
		n = s2.pop()
	}

	return n
}

var s1 = &stack{}
var s2 = &stack{}

func main() {
	for i := 0; i < 5; i++ {
		push(i)
	}

	fmt.Println(pop()) //0
	fmt.Println(pop()) //1

	push(6)
	push(7)

	fmt.Println(pop()) //2
	fmt.Println(pop()) //3
	fmt.Println(pop()) //4
	fmt.Println(pop()) //6

}
