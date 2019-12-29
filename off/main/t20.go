package main

import "fmt"

/*
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1））
*/

type stack1 struct {
	val    []int
	minNum []int
}

func (s *stack1) pop() {
	n := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	if n == s.min() {
		s.minNum = s.minNum[:len(s.minNum)-1]
	}
}

func (s *stack1) top() int {
	return s.val[len(s.val)-1]
}

func (s *stack1) push(n int) {
	s.val = append(s.val, n)
	if len(s.minNum) == 0 {
		s.minNum = append(s.minNum, n)
	} else if n <= s.min() {
		s.minNum = append(s.minNum, n)
	}
}

func (s *stack1) min() int {
	return s.minNum[len(s.minNum)-1]
}

func main() {
	s := stack1{}
	s.push(3)
	s.push(2)
	s.push(4)
	s.push(1)
	s.push(5)
	fmt.Println(s.min())
	s.pop()
	fmt.Println(s.min())
	s.pop()
	fmt.Println(s.min())

}
