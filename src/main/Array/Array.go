package main

import "fmt"

func main() {
	// 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成，一旦声明了，数组的长度就固定了，不能动态变化。
	//`len()` 和 `cap()` 返回结果始终一样。

	// 声明数组

	// 一维数组
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr4)

	arr5 := [5]int{0: 3, 1: 5, 4: 6} //意思是 0的位置是3 1的位置是5 4的位置是6
	fmt.Println(arr5)

	// 二维数组
	var arr6 = [3][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(arr6)

	arr7 := [3][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(arr7)

	arr8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr8)

	// 注意事项

	// 一、数组不可动态变化问题，一旦声明了，其长度就是固定的。
	var arr9 = [5]int{1, 2, 3, 4, 5}
	// arr9[5] = 6 运行会报错：invalid array index 5 (out of bounds for 5-element array)
	fmt.Println(arr9)

	// 二、数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销。
	var arr10 = [5]int{1, 2, 3, 4, 5}
	modifyArr(arr10)
	fmt.Println(arr10) // 不会改变原数组的值
	modifyArr1(&arr10)
	fmt.Println(arr10)

	// 三、数组赋值问题，同样类型的数组（长度一样且每个元素类型也一样）才可以相互赋值，反之不可以。
	var arr = [5]int{1, 2, 3, 4, 5}
	var arr11 [5]int = arr
	//var arr12 [6]int = arr 运行会报错：cannot use arr (type [5]int) as type [6]int in assignment
	fmt.Println(arr11)
}

func modifyArr(a [5]int) {
	a[1] = 20
}

func modifyArr1(a *[5]int) {
	a[1] = 20
}
