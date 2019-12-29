package main

import "fmt"

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，
序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。
（注意：这两个序列的长度是相等的）
*/

type stack2 struct {
	val []int
}

func (s *stack2) push(n int) {
	s.val = append(s.val, n)
}

func (s *stack2) pop() int {
	n := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return n
}

func (s *stack2) top() int {
	return s.val[len(s.val)-1]
}

// 首先将push数组压入栈中  在压入栈过程中如果栈顶元素等于pop数组元素头 那么就出栈  同时pop数组头往后移一位
// 如果最终栈为空则返回true  否则返回false
func IsPopOrder(push, pop []int) bool {
	s := stack2{}

	i := 0
	for _, v := range push {
		s.push(v)
		for s.top() == pop[i] {
			s.pop()
			i++
			if i == len(pop) {
				break
			}
		}
	}
	return len(s.val) == 0
}

func main() {
	push := []int{1,2,3,4,5}
	//pop := []int{5,4,3,2,1} // true
	//pop := []int{4,5,3,2,1}  // true
	//pop := []int{4,3,5,2,1}  // true
	pop := []int{4,3,5,1,2}  // false
	fmt.Println(IsPopOrder(push, pop))
}
