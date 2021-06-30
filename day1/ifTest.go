package main

import "fmt"

func main() {

	fmt.Print("有没有买西瓜的（Y/N）：")
	var yes string
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")
	if yes == "Y" || yes == "y" {
		fmt.Println("买一个西瓜")
	}

	fmt.Println("老公的想法：")
	if yes == "Y" || yes == "y" {
		fmt.Println("买一个西瓜和一个包子")
	} else {
		fmt.Println("买十个包子")
	}

	fmt.Println("请输入分数：")
	var score int8
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}
