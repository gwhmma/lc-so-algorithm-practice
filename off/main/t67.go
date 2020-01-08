package main

import (
	"fmt"
	"time"
)

/*
给你一根长度为n的绳子，请把绳子剪成整数长的m段（m、n都是整数，n>1并且m>1），
每段绳子的长度记为k[0],k[1],...,k[m]。请问k[0]xk[1]x...xk[m]可能的最大乘积是多少？
例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18
输入一个数n，意义见题面。（2 <= n <= 60）
*/

// n的最大截取乘积等于 小于它的数的乘积(比n小的多个数相加=n)
func cut(n int) int {
	var rs = make([]int, n)

	if n < 5 {
		return n
	}

	// 前4个最大都是它本身 所以先初始化它
	for i := 0; i < 4; i++ {
		rs[i] = i + 1
	}

	// 求n的最大值  那么先求出 1 到 n-1 的每一个的最大值
	// 从比它小的数字中 比较 2个相加等于它的数的最大值的积  它的最大乘积 就是小于它的2个相加等于它的数的最大值的积的最大值
	for i := 5; i <= n; i++ {
		max := 0
		for j := 0; j < i-1; j++ {
			// 如果所有数字都已经使用过 就结束  避免重复
			if j > len(rs)/2 {
				break
			}
			if rs[j]*rs[i-2-j] > max {
				max = rs[j] * rs[i-2-j]
			}
		}
		rs[i-1] = max
	}

	return rs[n-1]
}

func main() {
	s := time.Now()
	fmt.Println(cut(60))
	fmt.Println(time.Since(s))
}
