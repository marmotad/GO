package main

import (
	"fmt"
)

func main() {
	var age int = 20
	fmt.Printf("%T %d\n", age, age)
	fmt.Println(0777) //八进制
	fmt.Println(0x10) //十六进制

	fmt.Println(1 + 2)
	fmt.Println(2 - 1)
	fmt.Println(1 * 3)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)

	age ++ //自增
	fmt.Println(age)
	age -- //自减
	fmt.Println(age)

	fmt.Println(2 > 4)
	fmt.Println(1 < 3)
	fmt.Println(2 >= 2)
	fmt.Println(2 <= 3)
	fmt.Println(2 == 2)
	fmt.Println(1 != 3)

	//赋值(=、+=、-=、*=、/=、%=、&=、|=、^=、<<=、>>=)
	age = 10
	age += 3 // age  = age + 3
	fmt.Println(age)
	age -= 3 // age = age - 3
	fmt.Println(age)
	age *= 2 // age = age * 2
	fmt.Println(age)
	age %= 2 // age = age % 2
	fmt.Println(age)
	age /= 3 // age = age / 3
	fmt.Println(age)
	var intA int = 10
	var uintB uint = 2
	fmt.Println(intA + int(uintB)) //强制类型转换，将uint类型转换成int类型
	fmt.Println(uint(intA) + uintB) //强制类型转换，将int类型转换成uint类型

	//byte, rune
	var a byte = 'A'
	var b rune = '中'
	fmt.Println(a, b)

	fmt.Printf("%T %c \n", a, a)
	fmt.Printf("%T %q\n", b, b)
}
