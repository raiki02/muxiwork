package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		n := 0
		count := 0
		fmt.Print("Press the number you want the counter to execute：")
		_, err := fmt.Scanln(&n)
		if n <= 0 {
			fmt.Println("wrong number ! Press another number.")
		} else if err != nil {
			fmt.Println("n must be Int type ! ")
		} else {
			fmt.Println("You chose the number :", n)
			func() {
				for i := 0; i < n; i++ {
					time.Sleep(time.Second)
					count++
					fmt.Println("count = ", count)
				}

			}()

		}

	}

}
