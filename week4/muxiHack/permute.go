package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	for i := 0; i <= re(n)-1; i++ {
		var a = make([]int, n)
		res = append(res, a)
	}
	ReturnFunc(nums, res, 0)
	return res
}
func re(n int) int {
	num := 1
	for i := n; i >= 1; i-- {
		num = num * i
	}
	return num
}

func change(arr []int, i int, j int) {
	var mid int
	mid = arr[i]
	arr[i] = arr[j]
	arr[j] = mid
}

func ReturnFunc(array []int, num [][]int, p int) {
	var t int
	if p == len(array)-1 {
		copy(num[t], array)
		t++
	} else {
		for i := p; i <= len(array)-1; i++ {
			var arr = make([]int, len(array))
			copy(arr, array)
			change(arr, i, p)
			ReturnFunc(arr, num, p+1)
		}
	}
}

var n int

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	key := make([]int, n)

	res := permute(key)
	fmt.Println(res)
}
