package main

import "fmt"

func main() {
	fmt.Print("有卖西瓜的吗？（Y/N）:")
	var yes string
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")

	if yes != "y" && yes != "Y" {
		goto END //跳转到指定标签
	}
	fmt.Println("买一个西瓜")
END: //定义标签，一般标签使用大写字母+ ‘:’

	//goto替代for循环实现计算1到100的和
	result := 0
	i := 0
START: //开始位置标签
	if i > 100 {
		goto FOREND //如果i > 100 打印最后结果
	}
	result += i
	i++
	goto START
FOREND:
	fmt.Println(result)

BREAKEND:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j == 4 {
				break BREAKEND //跳到循环外
			}
			fmt.Println(i, j, i*j)
		}
	}
}
