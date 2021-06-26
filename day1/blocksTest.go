package main

import "fmt"

func main() {
	outer := 1
	inter := 2
	//声明inter
	{
		fmt.Println(outer)
		//子作用域可以引用父作用域的内容；父作用域不能使用子作用域的内容
	}
	{
		inter := 10
		//在子作用域中重新声明inter
		fmt.Println(inter)
	}
	fmt.Println(inter)
}
