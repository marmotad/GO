package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 10, 30, 5}
	sort.Ints(nums)
	//对int类型切片排序
	fmt.Println(nums)

	names := []string{"aa", "DD", "b"}
	sort.Strings(names)
	//对字符串排序
	fmt.Println(names)

	heights := []float64{1.2, 6.7, 7, 5.4}
	sort.Float64s(heights)
	//对float64类型排序
	fmt.Println(heights)
}
