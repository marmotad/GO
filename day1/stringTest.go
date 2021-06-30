package main

import (
	"fmt"
)

func main() {
	//"" ==> 可解释字符串
	//`` ==> 原生字符串
	var name = "ha\tha"
	var desc = `xi\txi`
	fmt.Println(name)
	fmt.Println(desc)
	//算数运算：+ （连接）
	fmt.Println("a" + "b")
	//字符串关系运算
	fmt.Println("ab" == "bb")
	fmt.Println("ab" != "bb")
	fmt.Println("ab" >= "bb")
	fmt.Println("ab" <= "bb")
	fmt.Println("ab" > "bb")
	fmt.Println("ab" < "bb")
	//赋值运算
	str := "ha"
	str += "ha"
	fmt.Println(str)
	//索引
	str1 := "xixixi"
	fmt.Printf("%T %c\n", str1[0], str1[0])

	//切片
	fmt.Printf("%T %v\n", str1[0:2], str1[0:2])
	//查看字符串长度
	fmt.Println(len(str1))
}
