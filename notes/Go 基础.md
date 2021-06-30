#  Go 基础

## 1 程序结构

### 1.1 基本结构

```go
package main  //以package声明开头，说 明源文件所属的包

import "fmt" //import导入依赖的包，其次为包级别的变量、常量、类型和函数的声明和赋值。

func main()  {
	//函数中可定义局部的变量、常量
	fmt.Println("Hellow Word")
}
```

### 1.2 编译

运行：
`go build -work -x -o helloworld.exe main.go`
`go run -work -x main.go`

## 2 变量

+ 变量是指对一块存储空间定义名称，通过名称对存储空间的内容进行访问或修改，使用var进行变量声明
+ 在函数中声明了变量之后必须要使用；如果在函数外声明变量可以不使用
+ 变量定义后，若没有赋值，变量会有一个零值，不同类型变量零值不同
+ 变量只能声明一次

### 2.1 变量声明

#### 2.1.1 声明变量

```go
	var me string //定义字符串类型的变量me
	me = "fan" //对变量赋值
	fmt.Println(me) //打印变量
```

#### 2.1.2 定义并赋值函数

#### 2.1.2 定义多个相同类型变量并赋值

```go
	var user, name string = "fan", "fan" //定义多个相同类型的变量
	fmt.Println(user, name )
```

#### 2.1.3 定义多个不同类型变量并赋值

**方法一：**

```go
	var (
		//定义多个不同类型变量并赋值
		age int = 20
		height float32 = 1.64
	)
```

**方法二：**

```go
	var ss, aa = "haha", 33 
```

#### 2.1.4 简短声明

+ 简短声明只能在函数内使用
+ 简短声明可以同时声明多个变量

```go
	isBoy := true //简短声明，使用时可以省略变量类型，go会推到函数类型，只能在函数内使用
	fmt.Println(isBoy)
```

#### 2.1.5 赋值

+ 可以通过赋值交换变量的值

```go
	b, bb := "b", "bb"
	b, bb =  bb, b //对变量重新赋值（讲话两个变量的值）
	fmt.Println(b, bb)
```

## 3 常量

+ 用于定义不可被修改的的值，需要在编译过程中进行计算，只能为基础的数据类型布尔、数值、字符串，使用const进行常量声明
+ 常量不能修改值
+ 常量在函数内定义时，可以不使用

### 3.1 定义常量

#### 3.1.1 定义一个常量

```go
	const Name string = "fan" //定义常量并赋值
	fmt.Println(Name)
```

#### 3.1.2 省略类型

```go
	const AGE = 10 //定义一个常量，省略常量类型
	fmt.Println(AGE)
```

#### 3.1.3 定义多个类型相同的常量

```go
	const A, B string = "A", "B" //定义多个类型相同的常量
	fmt.Println(A, B)
```

#### 3.1.4 定义多个类型不同的常量

```go
	const (
		//定义多个类型不同的常量
		C int = 1
		D string = "ss"
	)
```

#### 3.1.5 定义多个常量（省略类型）

```go
	const E, F = "EE", 22 //定义多个常量（省略类型）
	fmt.Println(E, F)
```

#### 3.1.5 定义多个常量（省略类型和值）

```go
	const (
		// 定义多个常量（省略类型和值）
		C1 int = 1
		C2
		C3
	)
	fmt.Println(C1, C2, C3)

----
运行结果:
1 1 1
```

### 3.2 枚举

+ go没有明确的枚举定义，但可以借助iota标识符实现一组自增量的值实现枚举

```go
	const (
		//枚举
		I1 int = iota
		I2
		I3
	)
	fmt.Println(I1, I2, I3)
----
运行结果:
0 1 2
```

## 4 作用域

+ 定义标识符可以使用的范围
+ 使用{}定义作用域的范围
+ 子作用域可以引用父作用域的内容；父作用域不能使用子作用域的内容

### 4.1 引用父作用域

```go
	outer :=1
	{
		fmt.Println(outer)
		//子作用域可以引用父作用域的内容；父作用域不能使用子作用域的内容
	}
```

### 4.2 重新对父作用域中存在标识符声明

```go
	inter := 2
	//声明inter
	{
		inter := 10
		//在子作用域中重新声明inter
		fmt.Println(inter)
	}
	fmt.Println(inter)
```

## 5 打印

### 5.1 格式化输出

#### 5.1.1 输出的格式

| 动 词 | 功 能                                      |
| ----- | ------------------------------------------ |
| %v    | 按值的本来值输出                           |
| %+v   | 在 %v 基础上，对结构体字段名和值进行展开   |
| %#v   | 输出 Go 语言语法格式的值                   |
| %T    | 输出 Go 语言语法格式的类型和值             |
| %%    | 输出 % 本体                                |
| %b    | 整型以二进制方式显示                       |
| %o    | 整型以八进制方式显示                       |
| %d    | 整型以十进制方式显示                       |
| %x    | 整型以十六进制方式显示                     |
| %X    | 整型以十六进制、字母大写方式显示           |
| %U    | Unicode 字符                               |
| %c    | 输出单个字符                               |
| %q    | 单引号围绕的字符字面值，由Go语法安全地转义 |
| %f    | 浮点数                                     |
| %p    | 指针，十六进制方式显示                     |
| %t    | bool类型占位符                             |

## 6 数据类型

### 6.1 bool类型

+ bool值用于表示真假
+ 标识符：bool
+ 可选择：true/false

#### 6.1.1 定义bool类型

```go
	var zero bool
	//查看bool类型的零值
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)
```

#### 6.1.2 bool的操作

##### 6.1.2.1 逻辑运算

+ 与 &&， 或 || ，非 ！

##### 6.1.2.2 真值表

```go
	fmt.Println("&& 与运算")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

------
运行结果：
&& 与运算
true
false
false
false
```

**（&& 与）真值表**

| X     | y     | 返回  |
| ----- | ----- | ----- |
| TRUE  | TRUE  | TRUE  |
| TRUE  | FALSE | FALSE |
| FALSE | TRUE  | FALSE |
| FALSE | FALSE | FALSE |

```go
	fmt.Println("|| 或运算")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

---
运行结果
|| 或运算
true
true
true
false
```

**（|| 或）真值表**

| X     | Y     | 返回  |
| ----- | ----- | ----- |
| TRUE  | TRUE  | TRUE  |
| TRUE  | FALSE | TRUE  |
| FALSE | TRUE  | TRUE  |
| FALSE | FALSE | FALSE |

```go
	fmt.Println("! 取反")
	fmt.Println(!true)
	fmt.Println(!false)
```

**（!非）真值表**

| 非   | X     | 返回  |
| ---- | ----- | ----- |
| !    | TRUE  | FALSE |
| !    | FALSE | TRUE  |

#### 6.1.2.2 关系运算(==、!=)

```go
	isBoy := true
	isGirl := false
	fmt.Println(isBoy == isGirl)
	fmt.Println(isBoy != isGirl)
```

### 6.2 数值类型

#### 6.2.1 整型

![image-20210627181747832](Go 基础.assets/image-20210627181747832.png)

![Go语言学习笔记_2021_06_27_180905](Go 基础.assets/Go语言学习笔记_2021_06_27_180905.png)

##### 6.2.1.1 定义整型变量

```go
	var age int = 20
	fmt.Printf("%T %d\n", age, age)
	fmt.Println(0777) //八进制
	fmt.Println(0x10) //十六进制
----
运行结果：
int 20
511
16
```

##### 6.2.1.2 算数运算

+ 运算符：：+、-、*、/、%、++、--
+ ++ 和 -- 所操作的对象必须是变量，且必须在变量后面

```go
	var age int = 20

	fmt.Println(1 + 2)
	fmt.Println(2 - 1)
	fmt.Println(1 * 3)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)

	age ++ //自增
	fmt.Println(age)
	age -- //自减
	fmt.Println(age)

---
运行结果：
3
1
3
4
1
21
20
```

##### 6.2.1.3 关系运算

+ 运算符：  >、>=、<、<=、==、!=

```go
	fmt.Println(2 > 4)
	fmt.Println(1 < 3)
	fmt.Println(2 >= 2)
	fmt.Println(2 <= 3)
	fmt.Println(2 == 2)
	fmt.Println(1 != 3)

---
运行结果：
false
true
true
true
true
true
```

##### 6.2.1.4 位运算

+ 运算符：&、|、^、<<、>>、&^

##### 6.2.1.5 赋值运算

+ 赋值运算符：=、+=、-=、*=、/=、%=、&=、|=、^=、<<=、>>=

```go
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

---运行结果：
13
10
20
0
0   
```

##### 6.2.1.6 类型转换

+ 不同类型间不能直接运算，需要进行类型转换
+ 从大往小转时，可能会导致溢出

```
	fmt.Println(intA + int(uintB)) //强制类型转换，将uint类型转换成int类型
	fmt.Println(uint(intA) + uintB) //强制类型转换，将int类型转换成uint类型
```

##### 6.2.1.7  rune 和 byte

###### 6.2.1.7.1 定义

```go
	//byte, rune
	var a byte = 'A'
	var b rune = '中'
	fmt.Println(a, b)

---
运行结果：
65 20013
```

###### 6.2.1.7.2 [] rune 和 [] byte 区别

**go源码中的定义：**

```go
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32
```

#### 6.2.2 浮点数

+ Go提供float32和float64两种浮点类型

+ float类型零值是0.0

##### 6.2.2.1 定义浮点型

```go
	var height float32 //= 1.68
	fmt.Println(height)
	fmt.Printf("%T %f\n",height, height)
```

##### 6.2.2.2 浮点数字面量

+ 十进制表示法：3.1415926
+ 科学记数法：1e-5(表示：1乘10的-5次方)，1.05E1（表示1.05乘10的1次方）

```go
	a1 := 1.05E1
	fmt.Println(a1)

---
运行结果：
10.5
```

##### 6.2.2.3 浮点型算数运算

+ 算数运算符：：+、-、*、/、++、--
  + 注意：针对/除数不能为0

+ 浮点数在存褚时，有一定精度损耗，使用运算时也会有精度损耗

```go
	//浮点数算数运算
	fmt.Println(1.1 + 1.2)
	fmt.Println(1.1 - 1.2)
	fmt.Println(1.1 * 1.2)
	fmt.Println(1.1 / 1.2)
	f1 := 1.1
	f1 ++
	fmt.Println(f1)
	f1 --
	fmt.Println(f1)

---
运行结果：
-0.1
1.32
0.9166666666666666
2.1
1.1
```

##### 6.2.2.3 关系运算

+ ：>、>=、<、<=
  + 浮点型不能进行==或!=比较，可选择使用两个浮点数的差在一定区间内则认为相等

```go
	//关系运算
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 <= 1.2)
	fmt.Println(1.1 >= 1.2)
---
运行结果：
false
true
true
false
```

##### 6.2.2.4 赋值运算符

+ 运算符：=、+=、-=、*=、/=

```go
	//赋值运算
	a1 += 0.2 //其他运算符和+=类似
	fmt.Println(a1)
```

##### 6.2.2.5 类型转换

```go
	fmt.Printf("%T %T", a1, float32(a1))

---
运行结果：
float64 float32
```

##### 6.2.2.6 浮点类型格式化输出

+ 使用fmt.Printf进行格式化参数输出，占位符:
  +  %f、%F：十进制表示法
    +  %n.mf表示最小占n个宽度并且保留m位小数
  + l %e、%E：科学记数法表示
  + l %g、%G：自动选择最紧凑的表示方法%e(%E)或%f(%F)

```go
	fmt.Println(a1)
	fmt.Printf("%11.3f", a1)

---
运行结果：
10.7
     10.700
```

### 6.3 字符串类型

#### 6.3.1 定义字符串

+ 可解析字符串：通过双引号(")来创建，不能包含多行，支持特殊字符转义序列
+  原生字符串：通过反引号(`)来创建，可包含多行，不支持特殊字符转义序列

```go
	//"" ==> 可解释字符串
	//`` ==> 原生字符串
	var name = "ha\tha"
	var desc = `xi\txi`
	fmt.Println(name)
	fmt.Println(desc)

---
运行结果：
ha	ha
xi\txi
```

#### 6.3.2 字符串操作

##### 6.3.2.1 算数运算符（+：连接字符串）

```go
	fmt.Println("a" + "b")
```

##### 6.3.2.2 字符串关系运算

```
	fmt.Println("ab" == "bb")
	fmt.Println("ab" != "bb")
	fmt.Println("ab" >= "bb")
	fmt.Println("ab" <= "bb")
	fmt.Println("ab" > "bb")
	fmt.Println("ab" < "bb")
	
---
运行结果：
false
true
false
true
false
true
```

##### 6.3.2.3 赋值运算

+ 赋值运算符：+=

```go
	//赋值运算
	str := "ha"
	str += "ha"
	fmt.Println(str)
	
---
运行结果：
haha
```

##### 6.3.2.3 索引（针对只包含ascii字符的字符串） 

```go
	str1 := "xixixi"
	fmt.Printf("%T %c\n", str1[0], str1[0])

---
运行结果：
uint8 x
```

##### 6.3.2.4 切片（针 对只包含ascii字符的字符串） 

```go
	//切片
	fmt.Printf("%T %v\n", str1[0:2], str1[0:2])
	//查看字符串长度
	fmt.Println(len(str1))

---
运行结果：
string xi
6
```

### 6.4 指针

```go
	// 指针
	var cc *int //定义指针类型变量
	fmt.Println(cc) //打印指针类型默认值,指针类型默认值是nil
	cc = &A //获取变量地址使用&
	C := &A
	fmt.Printf("%T %T\n", C, cc)
	fmt.Println(*cc) //打印内存地址中的值使用*
	*cc = 4 //修改内存地址中的值
	fmt.Println(A) //修改内存地址中的值，指向相同内存地址的变量的值也会被修改

---
运行结果：  
<nil>
*int *int
1
4
```

### 6.5 scan

```go
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Println("你的姓名的：", name)
```

## 7 流程控制

### 7.1 条件语句if

+ 当if 表达式结果为 true ，则执行 if 语句块内代码，否则依次从上到下判断 else if 表达式结果，若结果为 true 则执行对应语句块内代码并退出 if else if else 语句，若 i f 和else if 表达式均为 false ，则执行 else 语句块内代码

#### 7.1.1 if

+ 当if 表达式的结果为 t rue 则执行语句块 内代码

```go
	fmt.Print("有没有买西瓜的（Y/N）：")
	var yes string
	fmt.Scan(&yes)

	fmt.Println("老婆的想法：")
	fmt.Println("买十个包子")
	if yes == "Y" || yes == "y" {
		fmt.Println("买一个西瓜")
	}
```

#### 7.1.2 if else

+ 当if 表达式结果为 true ，则执行 if 语句块内代码，否则执行 else 语句块内代码

```go
	fmt.Println("老公的想法：")
	if yes == "Y" || yes == "y" {
		fmt.Println("买一个包子")
	} else {
		fmt.Println("买十个包子")
	}
```

#### 7.1.3 else if

+ 成绩评优: [90, 优秀 (A), [80, 90) => 良好 (B), [60, 80) => 及格 (C), [0, 60) =>不及格 (D) 

```go
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
```

## 7.4 选择语句switch

+ 对于选择语句可以有0 个或多个 case 语句， 最多有 1 个 default 语句 选择条件为 true的 case 语句 块 开始执行并退出，若所有条件为 false ，则执行 default 语句块 并退出 。可以通过 fallthrough 修改执行退出行为，继续执行下一条的 case 或 default 语句块

#### 7.4.1 switch case 单值

```go
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
```

#### 7.4.2 switch case 表达式

```go
	fmt.Print("请输入分数：")
	var score int8
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >=60 :
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
```

#### 7.4.3 初始化子语句

#### 7.4.4 fallthrough

+ switch cas e 默认执行 case 语句后退出，若 需要继续执行 下一个 case 语句块，可以在 case语句块中使用 fullthrough 进行声明

```go
	fmt.Print("请输入分数：")
	var score int8
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >=60 :
		fmt.Println("D")
		fallthrough
	default:
		fmt.Println("E")
	}

---
运行结果： 
请输入分数：61
D
E
```

### 7.5 for

#### 7.5.1 for循环

```go
	num := 0
	//初始化子语句; 条件子语句;后置子语句
	for i := 1 ; i <= 100 ; i ++ {
		num += i
	}
	fmt.Println(num)
```

#### 7.5.2 类while

```go
	num := 0
	i := 1
	for i <= 100 {
		num += i
		i ++
	}
	fmt.Println(num)
```

#### 7.5.3 无限循环

```go
	for {
		i ++
		fmt.Println(i)
	}
```

#### 7.5.4 for range

+ 用于遍历 可迭代对象中的每个元素，例如字符串， 数组，切片 ，映射，通道 等

```go
	desc := "我爱中国"
	for i, j := range desc {
		fmt.Printf("%d %T %q\n", i, j, j )
	}
}

---
运行结果：
0 int32 '我'
3 int32 '爱'
6 int32 '中'
9 int32 '国'
```

#### 7.5.5 break和continue

+ break : 循环结束，终止循环
+ continue: 跳过本次循环，开始下次循环

```go
	fmt.Println("break")
	for i := 0 ; i < 5 ; i++ {
		if i == 3 {
			break //循环结束，终止循环
		} else {
		fmt.Println(i)
		}
	}
	fmt.Println("continue")
	for i := 0 ; i < 5 ; i++ {
		if i == 3 {
			continue //跳过本次循环，开始下次循环
		} else {
			fmt.Println(i)
		}
	}
	
---
运行结结果：
break
0
1
2
continue
0
1
2
4
```

 7.6 goto与lable

```go
	//goto替代for循环实现计算1到100的和
	result := 0
	i := 0
START: //开始位置标签
 	if i > 100 {
 		goto FOREND //如果i > 100 打印最后结果
	}
	result += i
	i ++
	goto START
	FOREND:
		fmt.Println(result)
```

+ break 和 continue 后也可以指定 label 用于指定跳出 或跳过 指定 label 同层 级的循环

```go
BREAKEND:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++{
			if i * j == 4 {
				break BREAKEND //跳到循环外
			}
			fmt.Println(i, j, i * j)
		}
	}

---
运行结果：
0 0 0
0 1 0
0 2 0
1 0 0
1 1 1
1 2 2
2 0 0
2 1 2
```









