package main

import (
	"fmt"
)

func main() {
	s := make([]int, 1, 1)
	old_s := cap(s)
	for i := 0; i < 10000; i++ {
		s = append(s, i*10)
		if old_s != cap(s) {
			fmt.Printf("old_s_cap(%d)-->new_s_cap(%d) \n", old_s, cap(s))
			old_s = cap(s)
		}
	}

	fmt.Println("----------------------")
	m := make(map[int]int, 1)
	old_m := len(m)
	for i := 0; i < 50; i++ {
		m[i] = i * 100
		if old_m != len(m) {
			fmt.Printf("old_m_cap(%d)-->new_m_cap(%d) \n", old_m, len(m))
			old_m = len(m)
		}
	}
}
