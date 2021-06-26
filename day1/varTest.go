package main

import "fmt"

func main() {
	var me string                        //定义字符串类型的变量me
	me = "fan"                           //对变量赋值
	fmt.Println(me)                      //打印变量
	var user, name string = "fan", "fan" //定义多个相同类型的变量
	fmt.Println(user, name)

	var (
		//定义多个不同类型变量并赋值
		age    int     = 20
		height float32 = 1.64
	)
	fmt.Println(age, height)

	var ss, aa = "haha", 33
	fmt.Println(ss, aa)

	isBoy := true //简短声明，使用时可以省略变量类型，go会推到函数类型，只能在函数内使用
	fmt.Println(isBoy)
	b, bb := "b", "bb"
	b, bb = bb, b //对变量重新赋值（讲话两个变量的值）
	fmt.Println(b, bb)
}
