package main

import (
	"fmt"
)

func main() {
	var scores map[string]int //nil映射

	fmt.Printf("%T %#v\n", scores, scores)

	fmt.Println(scores == nil)
	//初始化映射
	//通过字面量初始化
	scores = map[string]int{"A": 1, "B": 2}
	fmt.Println(scores)
	//make 函数初始化
	//初始化一个空的映射
	scores = make(map[string]int)
	fmt.Println(scores)
	//映射的操作
	// 通过key查询
	scores = map[string]int{"A": 1, "B": 2}
	fmt.Println(scores["A"])
	//访问不存在的key时，会返回对应value的零值，int零值是0，所以通过key访问不准确
	fmt.Println(scores["C"])
	//通过两个变量查询key是否存在
	v, yes := scores["A"]
	fmt.Println(yes)
	if yes {
		fmt.Println(v)
	}
	if v, ok := scores["C"]; ok {
		//通过if初始化子语句判断
		fmt.Println(v)
	}
	//修改和添加
	//当key存在时，为修改，不存在时为添加
	scores["A"] = 8
	fmt.Println(scores["A"])
	scores["C"] = 1
	fmt.Println(scores["C"])
	//删除
	delete(scores, "A")
	fmt.Println("A")
	scores["a"] = 8
	fmt.Println(scores)
	//获取当前映射的元素数
	fmt.Println(len(scores))
	//遍历映射
	for k, v := range scores {
		fmt.Println(k, ":", v)
	}

	var names map[string]map[string]string
	names = map[string]map[string]string{"aa": {"地方": "中国", "年龄": "22"}}
	fmt.Println(names)
	//增加元素
	names["bb"] = map[string]string{"地方": "中国", "年龄": "22"}
	fmt.Println(names)
	//删除元素
	delete(names, "bb") // map[string]string]string{"地方": "中国", "年龄": "22"}
	fmt.Println(names)
	//遍历映射
	for k, v := range names {
		fmt.Println(k, ":", v)
	}
	//修改
	names["aa"] = map[string]string{"地方": "中国", "年龄": "24"}
	fmt.Println(names)

}
