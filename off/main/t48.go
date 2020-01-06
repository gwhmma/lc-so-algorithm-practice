package main

import "fmt"

/*
求1+2+3+...+n，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
*/

func sum(n int) int {
	rs := n
	//_ = rs > 0 && ((rs += sum(n-1)) > 0)   // fuck !
	return rs
}

func main() {
	fmt.Println(sum(4))
}
