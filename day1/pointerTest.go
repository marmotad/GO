package main

import "fmt"

func main() {
	A := 1
	B := A
	B = 3
	fmt.Println(A, B)

	// 指针
	var cc *int     //定义指针类型变量
	fmt.Println(cc) //打印指针类型默认值,指针类型默认值是nil
	cc = &A         //获取变量地址使用&
	C := &A
	fmt.Printf("%T %T\n", C, cc)
	fmt.Println(*cc) //打印内存地址中的值使用*
	*cc = 4          //修改内存地址中的值
	fmt.Println(A)   //修改内存地址中的值，指向相同内存地址的变量的值也会被修改
}
