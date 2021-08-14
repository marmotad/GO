package main

import (
	"fmt"
)

func main() {
	var slice []int //定义切片
	fmt.Printf("%T %#v\n", slice, slice)
	fmt.Println(slice == nil)
	//在切片声明后，会被初始化为nil，表示暂不存在的切片
	fmt.Printf("%#v %d %d\n", slice, len(slice), cap(slice))
	//切片字面量赋值
	slice = []int{1, 2, 3}
	fmt.Printf("%#v\n", slice)
	fmt.Printf("%#v %d %d\n", slice, len(slice), cap(slice))

	//数组切片赋值
	arr := [3]int{1, 2, 3}
	fmt.Printf("%#v\n", arr)
	slice = arr[0:3]
	fmt.Printf("%#v\n", slice)
	//
	fmt.Printf("%#v %d %d\n", slice, len(slice), cap(slice))
	//make 函数
	slice = make([]int, 3)
	// make函数创建切片，此时需要传入一个参数来指定切片的长度,若只指定长度，则长度和容量默认相等
	fmt.Printf("%#v %d %d\n", slice, len(slice), cap(slice))
	slice = make([]int, 3, 5)
	//分别指定长度和容量，容量要大于等于长度
	fmt.Printf("%#v %d %d\n", slice, len(slice), cap(slice))
	//查询
	fmt.Println(slice[1])
	//遍历切片
	for _, j := range slice {
		fmt.Println(j)
	}
	//切片操作
	fmt.Printf("%T %v\n", slice[0:2], slice[0:2])
	fmt.Printf("%T %v\n", slice[0:2:4], slice[0:2:4])
	//当指定切片容量时，新切片容量为new_cap - start
	nums := slice[1:3:4]
	fmt.Printf("%#v %d %d %d\n", nums, nums, len(nums), cap(nums))
	//当未指定切片容量时，新切片容量为old _cap - start
	nums = slice[1:3]
	fmt.Printf("%#v %d %d %d\n", nums, nums, len(nums), cap(nums))

	//修改
	slice[1] = 10
	fmt.Println(slice[1])
	//增加
	slice = append(slice, 11)
	fmt.Printf("%#v %d %d\n", slice, len(slice), cap(slice))

	//删除
	//copy
	nums01 := []int{1, 2, 3, 4}
	nums02 := []int{10, 20, 30, 40, 50}
	//将nums01 copy 到 nums02
	copy(nums02, nums01)
	fmt.Println(nums01, nums02)

	//将nums02 copy 到 nums01
	nums02 = []int{10, 20, 30, 40, 50}
	copy(nums01, nums02)
	//由于nums02比numbs01大，所以只copy nums02到nums01时只copy前面的元素
	fmt.Println(nums01, nums02)

	//删除索引为0的和索引为1的
	fmt.Println(nums01, len(nums01))
	nums01 = nums01[0 : len(nums01)-1]
	fmt.Println(nums01)
	nums01 = nums01[1:]
	fmt.Println(nums01)
	//删除中间的元素
	fmt.Println(nums02)
	copy(nums02[2:], nums02[3:])
	//将要删除的元素用后面的元素覆盖
	fmt.Println(nums02)
	nums02 = nums02[0 : len(nums02)-1]
	//删除最后一位元素（因为中间的元素已被覆盖，后面的元素为重复的）
	fmt.Println(nums02)

	//队列和堆栈
	//队列：先进先出
	queue := []int{1, 2, 3, 4}
	queue = queue[1:]
	fmt.Println(queue)
	queue = queue[1:]
	fmt.Println(queue)
	//堆栈：先进后出
	stack := []int{1, 2, 3, 4}
	stack = stack[0 : len(stack)-1]
	fmt.Println(stack)
	stack = stack[0 : len(stack)-1]
	fmt.Println(stack)

	//多维切片
	//多维切片定义
	points := [][]int{}
	fmt.Printf("%T\n", points)
	points02 := make([][]int, 0)
	fmt.Printf("%T\n", points02)

	//向多维切片添加元素
	points = append(points, []int{1, 2, 3})
	fmt.Println(points)
	points = append(points, []int{10, 20, 30, 40, 50})
	//多维切片内的切片长度可以不同
	fmt.Println(points)

}
