package main

import (
	"fmt"
)

func main() {
	num := 0
	//初始化子语句; 条件子语句;后置子语句
	for i := 1; i <= 100; i++ {
		num += i
	}
	fmt.Println(num)

	num = 0
	i := 1
	for i <= 100 {
		num += i
		i++
	}
	fmt.Println(num)
	/*	for {
			i ++
			fmt.Println(i)
		}
	*/
	desc := "我爱中国"
	for i, j := range desc {
		fmt.Printf("%d %T %q\n", i, j, j)
	}
}
