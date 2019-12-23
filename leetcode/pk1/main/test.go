package main

import (
	"fmt"
	"sort"
)

func main()  {
	//s := "abcde"
	//for i := len(s); i > 0; i-- {
	//
	//	fmt.Println(s[i-1:i])
	//}

	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i:i+1])
	//}
	//fmt.Println(s[1:1])

	//s += "f"
	//fmt.Println(s)

	s := []int{2,1,7,3,5,8,4}
	sort.Ints(s)
	fmt.Println(s)
}
