package main

import (
	"fmt"
)

func main() {
	var zero bool
	//查看bool类型的零值
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)
	fmt.Println("&& 与运算")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println("|| 或运算")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println("! 取反")
	fmt.Println(!true)
	fmt.Println(!false)

	fmt.Println(isBoy == isGirl)
	fmt.Println(isBoy != isGirl)
	fmt.Printf("%T\n", isGirl)
	fmt.Printf("%t \n", isGirl)

}
