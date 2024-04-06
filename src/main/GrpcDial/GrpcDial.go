package main

import (
	"GoBeginner/src/main/friend"
	"fmt"
)

// conn, err := grpc.Dial("127.0.0.1:8000",
//
//		grpc.WithChainStreamInterceptor(),
//		grpc.WithInsecure(),
//		grpc.WithBlock(),
//		grpc.WithDisableRetry(),
//	)
//
// 咱们怎么能写出类似这样的调用方式，它是怎么实现的？
// 这篇文章咱们写一个 Demo，其实很简单，一步步往下看。
// ## 一
// `opts …DialOption`，这个是不定参数传递，参数的类型为 `DialOption`，不定参数是指函数传入的参数个数为不定数量，可以不传，也可以为多个。
// 写一个不定参数传递的方法也很简单，看看下面这个方法 1 + 2 + 3 = 6。
func Add(a int, args ...int) (result int) {
	result += a
	for _, arg := range args {
		result += arg
	}
	return result
}

func main() {
	fmt.Println(Add(1, 2, 3))
	//其实平时我们用的 `fmt.Println()`、`fmt.Sprintf()` 都属于不定参数的传递。

	//## 二
	//`WithInsecure()`、`WithBlock()` 类似于这样的 With 方法，
	//其实作用就是修改 `dialOptions` 结构体的配置，之所以这样写我个人认为是面向对象的思想，当配置项调整的时候调用方无需修改。
	//## 场景
	//咱们模拟一个场景，使用 `不定参数` 和 `WithXXX` 这样的写法，写个 Demo，
	//比如我们要做一个从附近找朋友的功能，配置项有：性别、年龄、身高、体重、爱好，我们要找性别为女性，
	//年龄为30岁，身高为160cm，体重为55kg，爱好为爬山的人，希望是这样的调用方式：
	friends, err := friend.Find("附近的人",
		friend.WithSex(1),
		friend.WithAge(30),
		friend.WithHeight(160),
		friend.WithWeight(55),
		friend.WithHobby("爬山"))

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(friends)

}
