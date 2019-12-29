package main

import "math"

/*
给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
保证base和exponent不同时为0
*/

func Power1(base float64, n int64) float64 {
	return math.Pow(base, float64(n))
}

func Power2(base float64, n int) float64 {
	if base == 0 {
		return 0
	} else if n == 0 {
		return 1.0
	}

	abs := 0
	res := 1.0
	if n > 0 {
		abs = n
	} else {
		abs = -n
	}

	for i := 0; i < abs; i++ {
		res *= base
	}

	if n > 0 {
		return res
	}
	return 1 / res
}

func main() {

}
