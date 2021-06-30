package main

import "fmt"

func main() {
	fmt.Println("break")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break //循环结束，终止循环
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("continue")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue //跳过本次循环，开始下次循环
		} else {
			fmt.Println(i)
		}
	}
}
