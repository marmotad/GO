package main

import (
	"fmt"
)

func main() {
	var height float32 //= 1.68
	fmt.Println(height)
	fmt.Printf("%T %f\n", height, height)

	a1 := 1.05e1
	fmt.Println(a1)
	//浮点数算数运算
	fmt.Println(1.1 + 1.2)
	fmt.Println(1.1 - 1.2)
	fmt.Println(1.1 * 1.2)
	fmt.Println(1.1 / 1.2)
	f1 := 1.1
	f1++
	fmt.Println(f1)
	f1--
	fmt.Println(f1)
	//关系运算
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 <= 1.2)
	fmt.Println(1.1 >= 1.2)

	//赋值运算
	a1 += 0.2 //其他运算符和+=类似
	fmt.Println(a1)

	fmt.Printf("%T %T\n", a1, float32(a1))
	fmt.Println(a1)
	fmt.Printf("%11.3f", a1)
}
