package main

import "fmt"

func main() {
	//投票计数
	users := []string{"a", "b", "a"}
	//方法一
	//投票序列
	votes := map[string]int{}
	//使用映射统计票数
	for _, v := range users {
		//遍历数组中成员
		if _, ok := votes[v]; ok {
			//如果当前数组成员在映射中存在，映射的value值加一
			votes[v] = votes[v] + 1
		} else {
			//如果当前数组成员在映射中不存在，在映射中添加该成员并赋值为1
			votes[v] = 1
		}
	}
	fmt.Println(votes)

	//方法二
	votes = map[string]int{}
	//使用映射统计票数
	for _, v := range users {
		//遍历数组中成员
		votes[v]++
		//对遍历到的key对应的value值加1（由于映射遍历到的不存在的key值时，默认value值为0）
	}
	fmt.Println(votes)
}
