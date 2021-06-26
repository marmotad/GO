package main

import "fmt"

func main() {
	const NAME string = "fan" //定义常量并赋值
	fmt.Println(NAME)

	const AGE = 10 //定义一个常量，省略常量类型
	fmt.Println(AGE)
	const A, B string = "A", "B" //定义多个类型相同的常量
	fmt.Println(A, B)
	const (
		//定义多个类型不同的常量
		C int    = 1
		D string = "ss"
	)
	fmt.Println(C, D)
	const E, F = "EE", 22 //定义多个常量（省略类型）
	fmt.Println(E, F)
	const (
		// 定义多个常量（省略类型和值）
		C1 int = 1
		C2
		C3
	)
	fmt.Println(C1, C2, C3)

	const (
		//枚举
		I1 int = iota
		I2
		I3
	)
	fmt.Println(I1, I2, I3)
	//
}
