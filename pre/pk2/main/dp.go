package main

import "fmt"

//钢条切割
func cut(price []int, n int) int {
	var s = make([]int, n+1)
	i := 1

	for i <= n {
		j := 1
		tmp := 0
		for j <= i {
			tmp = max(tmp, price[j-1]+s[i-j])
			j++
		}
		s[i] = tmp
		i++
	}
	return s[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	p := []int{1,5,8,9,10,17,17}
	fmt.Println(cut(p,5))
}
