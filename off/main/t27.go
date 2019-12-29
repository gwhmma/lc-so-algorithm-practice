package main

import (
	"fmt"
	"sort"
)

/*
输入一个字符串,按字典序打印出该字符串中字符的所有排列。
例如输入字符串abc,则打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
*/

type stack5 struct {
	val []int
}

func (s *stack5) push(n int) {
	s.val = append(s.val, n)
}

func (s *stack5) pop() int {
	n := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return n
}

func (s *stack5) all() []int {
	return s.val
}

//回溯法
func Permutation(s string) []string {
	var res []string
	var st stack5
	used := make([]bool, len(s))
	var data []int

	if len(s) == 0 {
		return res
	}

	bte := []byte(s)
	for _, v := range bte {
		data = append(data, int(v))
	}
	sort.Ints(data)

	var dfs func(data []int, depth int, s stack5)
	dfs = func(data []int, depth int, s stack5) {
		if depth == len(data) {
			tmp := ""
			all := s.all()
			for _, v := range all {
				tmp += string(v)
			}
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(data); i++ {
			if !used[i] {
				if i > 0 && data[i] == data[i-1] && !used[i-1] {
					continue
				}

				s.push(data[i])
				used[i] = true
				dfs(data, depth+1, s)
				s.pop()
				used[i] = false
			}
		}
	}

	dfs(data, 0, st)
	return res
}

func main() {
	s := "aabcxx"
	fmt.Println(Permutation(s))
}
