package main

import (
	"fmt"
)

func main() {
	s := "hello,世界"
	len := len(s)
	fmt.Println("len(s) = ", len)

	for i, _ := range s {
		b := string(s[i])
		fmt.Print(" ", b)
	}

	fmt.Println(" \n ---")
	for i := 0; i < len; i++ {
		c := string(s[i])
		fmt.Print(" ", c)
	}
	fmt.Println(" \n ---")
	for _, i := range s {
		d := string(i)
		fmt.Print(" ", d)
	}
}
