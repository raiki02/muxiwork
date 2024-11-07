/* 使用for循环生成20个goroutine，并向一个channel传入随机数和goroutine编号，等待这些goroutine都生成完后，
想办法给这些goroutine按照编号进行排序(输出排序前和排序后的结果,要求不使用额外的空间存储着20个数据)*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	order := make(chan int, 20)
	for i := 0; i < 20; i++ {
		r := rand.Intn(100)
		order <- r
		fmt.Print(" ", r)

		go getOrder(order)
	}
}
