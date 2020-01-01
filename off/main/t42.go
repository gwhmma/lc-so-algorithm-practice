package main

import (
	"fmt"
	"time"
)

/*
小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,他马上就写出了正确答案是100。
但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。
没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。
现在把问题交给你,你能不能也很快的找出所有和为S的连续正数序列? Good Luck!
*/

type List1 struct {
	list []int
}

// 等差数列  公差 d = 1   第一项 a1    n : 有多少个数
// s = n(2a1 + n -1) / 2
// 得到 s 的因子   如 10 = 1 * 10  2 * 5
func sequence(s int) *[]List1 {
	var rs []List1
	var factor = make(map[int]int)
	fac(2*s, &factor)

	for k, v := range factor {
		n := k
		a1 := (v - n + 1) / 2

		if a1 <= 0 || a1 == s {
			continue
		}

		res := a1
		tmp := a1
		var arr []int
		arr = append(arr, res)
		for i := 1; i < n; i++ {
			tmp += 1
			arr = append(arr, tmp)
			res += tmp

		}
		if res == s {
			list := List1{list: arr}
			rs = append(rs, list)
		}
	}
	return &rs
}

// 得到 s 的所有因子
func fac(s int, m *map[int]int) {
	for i := 1; i <= s; i++ {
		if i*(s/i) == s {
			(*m)[i] = s / i
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println(sequence(300000000))
	end := time.Since(start)
	fmt.Println(end)
}
