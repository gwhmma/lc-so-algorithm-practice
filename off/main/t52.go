package main

import "fmt"

/*
给定一个数组A[0,1,...,n-1],请构建一个数组B[0,1,...,n-1],
其中B中的元素B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]。
不能使用除法
*/

/*
忽略 0 乘以任意数等于0
a = [0,1,2,3,4]

  0 1 2 3 4
0 '0' 1 2 3 4
1 0 '1' 2 3 4
2 0 1 '2' 3 4
3 0 1 2 '3' 4
4 0 1 2 3 '4'

----->
0
0 1
0 1 2
0 1 2 3

1 2 3 4
  2 3 4
    3 4
      4
*/

// 可将整个求值过程分解为以上2个递增的三角形
// 2个三角形的积即为对应数组下标的值
func multiply(a []int) []int {
	var rs = make([]int, len(a))
	var down = make([]int, len(a)-1)
	var up = make([]int, len(a)-1)
	down[0] = a[0]
	up[len(up)-1] = a[len(a)-1]

	for i := 1; i < len(a)-1; i++ {
		down[i] = down[i-1] * a[i]
	}
	for i := len(a) - 2; i > 0; i-- {
		up[i-1] = up[i] * a[i]
	}

	rs[0] = up[0]
	rs[len(rs)-1] = down[len(down)-1]

	for i := 0; i < len(up)-1; i++ {
		rs[i+1] = down[i] * up[i+1]
	}

	return rs
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(multiply(a))
}
