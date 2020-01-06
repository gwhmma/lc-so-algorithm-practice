package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "12!"
	s1 := strings.Split(s, "!")
	fmt.Println(len(s1))

}
