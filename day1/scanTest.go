package main

import "fmt"

func main() {
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Println("你的姓名的：", name)
}
