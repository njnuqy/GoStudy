package main

import "fmt"

func main() {
	// 数据类型
	// 1.字符串 只能用一对双引号（""）或反引号（``）括起来定义，不能用单引号（''）定义！
	const name1 string = "Tom"
	const name2 string = `Jerry`
	// 2. 整形  `int8` `uint8`
	//`int16` `uint16`
	//`int32` `uint32`
	//`int64` `uint64`
	//`int` `uint`，具体长度取决于 CPU 位数。
	//**浮点型**
	//`float32` `float64`
	const age1 float32 = 12.2
	const age2 int = 13
	fmt.Println(name1+"年龄为：", age1)
	fmt.Println(name2, age2)

	// 变量声明
	// 第一种：var 变量名称 数据类型 = 变量值，如果不赋值，使用的是该数据类型的默认值。
	var variable1 int = 1
	// 第二种：var 变量名称 = 变量值 , 根据变量值，自行判断数据类型。
	var variable2 = 2
	// 第三种：变量名称 := 变量值 , 省略了 var 和数据类型，变量名称一定要是未声明过的。
	varibale3 := 3
	fmt.Println(variable1, variable2, varibale3)
	// 多个变量声明
	// 第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...
	var mutipleVariable1, mutipleVariable2 int = 4, 5
	// 第二种：var 变量名称,变量名称 ...  = 变量值,变量值 ...
	var mutipleVariable3, mutipleVariable4 = 6, 7
	// 第三种：变量名称,变量名称 ... := 变量值,变量值 ...
	mutipleVariable5, mutipleVariable6 := 8, 9
	fmt.Println(mutipleVariable1, mutipleVariable2, mutipleVariable3, mutipleVariable4, mutipleVariable5, mutipleVariable6)

	// ## 输出方法
	//**fmt.Print**：输出到控制台（仅只是输出）
	//**fmt.Println**：输出到控制台并换行
	//**fmt.Printf**：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）
	//**fmt.Sprintf**：格式化并返回一个字符串，不输出。
	fmt.Print("输出到控制台不换行")
	fmt.Println("---")
	fmt.Println("输出到控制台并换行")
	fmt.Printf("name=%s,age=%d\n", "Tom", 30)
	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}
