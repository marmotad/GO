package main

import "fmt"

func main() {

	fmt.Print("有没有买西瓜的（Y/N）：")
	var yes string
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")

	switch yes {
	case "Y", "y":
		fmt.Println("买一个西瓜")
	}

	fmt.Println("老公的想法")
	switch yes {
	case "Y", "y":
		fmt.Println("买一个包子")
	default:
		fmt.Println("买十个包子")
	}

	fmt.Print("请输入分数：")
	var score int8
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 60:
		fmt.Println("D")
		fallthrough
	default:
		fmt.Println("E")
	}
}
