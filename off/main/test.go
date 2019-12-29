package main

import "fmt"

func main() {
	a := []int{97,98,99}
	s := ""
	i := 0
	for i < len(a) {
		s += string(a[i])
		i++
	}

	fmt.Println(s)

}
