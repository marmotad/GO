package main

import "fmt"

func main() {
	var nums [10]int
	var bools [3]bool
	//声明数组，数组的零值是对应数据类型的零值

	fmt.Printf("%T\n ", nums)
	fmt.Println(nums)
	fmt.Println(bools)

	//声明指定长度数组
	nums1 := [3]int{1, 2, 3}
	fmt.Println(nums1)
	//声明指定长度数组，使用索引赋值
	nums2 := [3]int{0: 11, 2: 20}
	fmt.Println(nums2)
	//使用推导方式声明数组
	nums3 := [...]int{10, 20, 30}
	fmt.Println(nums3)
	//使用推导方式声明数组,并使用索引赋值
	nums4 := [...]int{0: 10, 2: 20}
	fmt.Println(nums4)

	//字面量赋值
	//指定数组长度赋值
	nums = [10]int{1, 2, 3}
	fmt.Println(nums)
	//对指定位置元素进行初始化
	nums = [10]int{0: 10, 1: 20, 9: 10}
	fmt.Println(nums)
	//使用初始化元素数量推到数组长度
	nums = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)

	//关系运算
	nums5 := [3]int{1, 2, 3}
	nums6 := [3]int{1, 2, 4}
	fmt.Println(nums5 == nums6)
	fmt.Println(nums5 != nums6)
	//获取数组长度
	fmt.Println(len(nums5))
	//数组索引
	//使用索引打印数组中元素
	fmt.Println(nums5[0])
	//通过索引修改数组中的值
	nums5[1] = 100
	fmt.Println(nums5[1])
	//遍历数组中的数据
	//普通for遍历
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	//for range遍历
	for _, value := range nums {
		fmt.Println(value)
	}
	//数组的切片
	//array[start:end]
	fmt.Printf("%T\n", nums[1:3]) //数组的切片是切片类型
	fmt.Printf("%v\n", nums[1:3]) //打印切片数组

	//array[start:end:cap](end<=cap<=len)获取数组的一部分元素做为切片
	fmt.Printf("%T\n", nums[0:3:3])
	fmt.Printf("%v\n", nums[0:3:4])
	//多维数组
	var marrays [3][3]int
	fmt.Println(marrays)
	//多维数组赋值
	marrays[1] = [3]int{1, 2, 3}
	fmt.Println(marrays)
	marrays[0][1] = 10
	fmt.Println(marrays)
}
