package main

import "fmt"

/*
go into ch(ns)
out -> oc


*/

func main() {
	ch_n := make(chan string, 2)
	ch_s := make(chan string, 2)
	s1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s2 := "0123456789101112131415161718192021222324252627282930"
	i1 := 0
	i2 := 0
	var oc string
	go func() {
		for i1 < len(s1) {
			ch_s <- string(s1[i1])
			i1++
			ch_s <- string(s1[i1])
			i1++

		}
	}()

	go func() {
		for i2 < len(s2) {
			ch_n <- string(s2[i2])
			i2++
			ch_n <- string(s2[i2])
			i2++

		}
	}()

	var temp_s string
	var temp_n string

	for i := 0; i < len(s2)/2; i++ {
		select {
		case temp_s = <-ch_s:
			oc += temp_s
		case temp_s = <-ch_s:
			oc += temp_s

		case temp_n = <-ch_n:
			oc += temp_n

		case temp_n = <-ch_n:
			oc += temp_n

		}
	}

	fmt.Println(oc)
}

// func main() {
// 	ch_n := make(chan byte, 2)
// 	ch_s := make(chan byte, 2)
// 	s1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	s2 := "0123456789101112131415161718192021222324252627282930"
// 	var oc string
// 	// for i := 0; i < len(s1); i++ {
// 	// 	fmt.Println(s1[i])
// 	// }
// 	go func() {
// 		for _, i := range s1 {
// 			ch_n <- i
// 		}
// 		close(ch_n)

// 	}()
// 	go func() {
// 		for {
// 			for _, i := range s2 {
// 				ch_s <- i
// 			}
// 		}
// 	}()
// 	for {
// 		i := 1
// 		if i%4 == 1 || i%4 == 2 {
// 			c := <-ch_n
// 			oc += string(c)
// 			if c == 'Z' {
// 				close(ch_s)
// 				break
// 			}
// 		} else {
// 			c := <-ch_s
// 			oc += string(c)
// 		}
// 		i++
// 	}

// }
