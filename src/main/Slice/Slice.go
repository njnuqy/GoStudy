package main

import "fmt"

func main() {
	// 切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。 `len()` 和 `cap()` 返回结果可相同和不同。

	// 声明切片
	var slice1 []int // nil切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice1), cap(slice1), slice1)

	var slice2 = []int{} // 空切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)

	var slice3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice3), cap(slice3), slice3)

	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice4), cap(slice4), slice4)

	var slice5 []int = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice5), cap(slice5), slice5)

	slice6 := make([]int, 6, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice6), cap(slice6), slice6)

	// 截取切片
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("slice[1] == ", slice[1])
	fmt.Println("slice[:] == ", slice[:])
	fmt.Println("slice[1:] == ", slice[1:])
	fmt.Println("slice[:4] == ", slice[:4])
	fmt.Println("slice[0:3] == ", slice[0:3])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice[0:3]), cap(slice[0:3]), slice[0:3])
	//在 Go 语言中，当你对一个切片（slice）使用三个索引进行切片操作时形式为 slice[low:high:max]，
	//这里的 low 是切片的起始索引（包含），high 是切片的结束索引（不包含），而 max 是切片的容量上限（capacity）
	fmt.Println("sli[0:3:4] ==", slice[0:3:4])
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice[0:3:4]), cap(slice[0:3:4]), slice[0:3:4])

	// 追加切片
	sli := []int{4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)
	// append 时，容量不够需要扩容时，cap 会翻倍。
	sli = append(sli, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	// 删除切片
	delsli := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(delsli), cap(delsli), delsli)

	//删除尾部 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(delsli[:len(delsli)-2]), cap(delsli[:len(delsli)-2]), delsli[:len(delsli)-2])

	//删除开头 2 个元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(delsli[2:]), cap(delsli[2:]), delsli[2:])

	//删除中间 2 个元素
	delsli = append(delsli[:3], delsli[3+2:]...)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(delsli), cap(delsli))
}
