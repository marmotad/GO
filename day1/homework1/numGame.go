package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var yes string
	var guess int
	//定义一个变量用于存储猜测的数
	start := true
	//定义一个变量初始值为true,当退出程序时，将变量值修改为false
	rand.Seed(time.Now().Unix())
	for start {
		num := rand.Int() % 100
		//生成一个随机数
		fmt.Println(num)
		for i := 1; i <= 5; i++ {
			fmt.Print("请输入你猜的数字：")
			fmt.Scan(&guess)
			if num == guess {
				//判断猜测的是否正确
				fmt.Println("猜对了！")
				fmt.Printf("你用了%d次机会！\n", 5-i)
				fmt.Print("是否再来一次（Y/N）:")
				fmt.Scan(&yes)
				if yes == "y" || yes == "Y" {
					//如果重来一次，退出本次循环，重新开始循环
					break
				} else {
					start = false
					//如果输入的不是y或Y,认为结束游戏，修改for循环条件为false，退出最外层循环
					break
				}
			} else if num < guess {
				fmt.Println("太大了！")
				fmt.Printf("你还有%d次机会！\n", 5-i)
			} else if num > guess {
				fmt.Println("太小了")
				fmt.Printf("你还有%d次机会！\n", 5-i)
			}
			if i == 5 {
				fmt.Print("是否再来一次（Y/N）:")
				fmt.Scan(&yes)
				if yes == "y" || yes == "Y" {
					break
					//如果重来一次，退出本次循环，重新开始循环
				} else {
					start = false
					//如果输入的不是y或Y,认为结束游戏，修改for循环条件为false，退出最外层循环
				}
			}
		}
	}
}
