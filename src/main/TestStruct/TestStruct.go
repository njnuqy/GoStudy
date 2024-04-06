package main

import (
	"GoBeginner/src/main/study"
	"fmt"
)

//在 Go 语言中，一个 struct 实现了某个接口里的所有方法，就叫做这个 struct 实现了该接口。
//下面写一个 Demo 实现一下，先写一个 Study interface{}
//，里面需要实现 4 个方法 Listen、Speak、Read、Write，然后再写一个 study struct{}，去全部实现里面的方法，然后分享一下代码心得

func main() {
	name := "Tom"
	s, err := study.New(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}

//## 代码解释
//#### 一、
//var _ Study = (*study)(nil)
//要求 `*study` 去实现 `Study`，若 `Study` 接口被更改或未全部实现时，在编译时就会报错。
//#### 二、
//type study struct {
//	Name string
//}
//之所以定义为私有的结构体，是因为不想在其他地方被使用，比如后面将 `Name` 改成 `UserName` 只需要在本包内修改即可。
//#### 三、
//```
//func New(name string) (Study, error) {
//	if name == "" {
//		return nil, errors.New("name required")
//	}
//	return &study{
//		Name: name,
//	}, nil
//}
//```
//在其他地方调用 `New()` 使用 `Study` 包时，仅对外暴露了 4 个方法，别人只管调用就好了，内部实现别人无需关心。
