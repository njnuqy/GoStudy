package main

import "fmt"

func main() {
	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n", len(person), cap(person), person)

	// 循环Array
	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}
	fmt.Println("")

	// 使用空白符
	for _, name := range person {
		fmt.Println("name :", name)
	}

	// 循环slice
	fmt.Println("================")
	slicePerson := []string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slicePerson), cap(slicePerson), slicePerson)
	fmt.Println("================")

	for k, v := range slicePerson {
		fmt.Printf("slicePerson[%d]: %s\n", k, v)
	}
	fmt.Println("================")

	for i := range slicePerson {
		fmt.Printf("slicePerson[%d]: %s\n", i, slicePerson[i])
	}
	fmt.Println("================")

	for i := 0; i < len(slicePerson); i++ {
		fmt.Printf("slicePerson[%d]: %s\n", i, slicePerson[i])
	}
	fmt.Println("================")

	// 使用空白符
	for _, name := range slicePerson {
		fmt.Println("name : ", name)
	}

	// 循环map
	mapPerson := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}
	fmt.Printf("len=%d map=%v\n", len(mapPerson), mapPerson)
	fmt.Println("================")

	for k, v := range mapPerson {
		fmt.Printf("mapPerson[%d]: %s\n", k, v)
	}
	fmt.Println("================")

	for i := range mapPerson {
		fmt.Printf("mapPerson[%d]: %s\n", i, mapPerson[i])
	}
	fmt.Println("================")

	for i := 0; i < len(mapPerson); i++ {
		fmt.Printf("mapPerson[%d]: %s\n", i, mapPerson[i])
	}
	fmt.Println("================")

	for _, name := range mapPerson {
		fmt.Println("name :", name)
	}

	// break 用于退出当前循环，可用于for,select,switch
	for i := 1; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println("i = ", i)
	}
	fmt.Println("================")
	// continue 跳过本次循环，只能用于for
	for i := 1; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println("i = ", i)
	}
	fmt.Println("================")
	// goto改变函数内代码执行顺序，不能夸函数使用

	for i := 1; i < 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i = ", i)
	}
END:
	fmt.Println("end")

	i := 3
	fmt.Printf("当 i = %d 时：\n", i)

	switch i {
	case 1:
		fmt.Println("i=", i)
	case 2:
		fmt.Println("i=", i)
	case 3:
		fmt.Println("i=", i)
		fallthrough
	case 4, 5, 6:
		fmt.Println("i=", "4,5,6")
	default:
		fmt.Println("i=", i)
	}
	//**结论：**
	//- 默认每个 case 带有 break
	//- case 中可以有多个选项
	//- fallthrough 不跳出，并执行下一个 case
}
