package main

import (
	"fmt"
	"os"
	"time"
)

// defer 函数大家肯定都用过，它在声明时不会立刻去执行，而是在函数 return 后去执行的。
// 它的主要应用场景有异常处理、记录日志、清理数据、释放资源 等等。
// 这篇文章不是分享 defer 的应用场景，而是分享使用 defer 需要注意的点。
// 咱们先从一道题开始，一起来感受下 ...
func main() {
	// 执行顺序

	// 例1
	var x int = 1
	var y int = 2
	defer calc("A", x, calc("B", x, y))
	x = 3
	defer calc("C", x, calc("D", x, y))
	y = 4

	// 例2
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("main")

	// 结论：defer 函数定义的顺序 与 实际执的行顺序是相反的，也就是最先声明的最后才执行。

	// 闭包
	var a int = 1
	var b int = 2
	defer fmt.Println(a + b)
	defer func() {
		fmt.Println(a + b)
	}()

	// 结论：闭包获取变量相当于引用传递，而非值传递。
	defer func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
	a = 2
	fmt.Println("main")
	// 结论：传参是值复制。
	//还可以理解为：defer 调用的函数，参数的值在 defer 定义时就确定了，看下代码
	//`defer fmt.Println(a + b)`，在这时，参数的值已经确定了。
	//而 defer 函数内部所使用的变量的值需要在这个函数运行时才确定，看下代码
	//`defer func() { fmt.Println(a + b) }()`，a 和 b 的值在函数运行时，才能确定。
	fmt.Println(t1()) // 1
	fmt.Println(t2()) // 2
	fmt.Println(t3()) // 1
	fmt.Println(t4()) // 1
	// 结论：return 不是原子操作。

	//os.Exit
	defer fmt.Println("1")
	fmt.Println("main")
	os.Exit(0)
	// 结论：当`os.Exit()`方法退出程序时，defer不会被执行。

	// 不同协程
	goA()
	time.Sleep(1 * time.Second)
	fmt.Println("main")
	// `GoB()` panic 捕获不到。
	//结论：defer 只对当前协程有效。
	//这个问题怎么解？咱们下回再说。
	//接下来，咱们分析下文章开头的问题吧。
	//## 答案解析
	//先列出答案：
	//B 1 2 3
	//D 3 2 5
	//C 3 5 8
	//A 1 3 4
	//其实上面那道题，可以拆解为：
	//func calc(index string, a, b int) int {
	//	ret := a + b
	//	fmt.Println(index, a, b, ret)
	//	return ret
	//}
	//func main() {
	//	x := 1
	//	y := 2
	//	tmp1 := calc("B", x, y)
	//	defer calc("A", x, tmp1)
	//	x = 3
	//	tmp2 := calc("D", x, y)
	//	defer calc("C", x, tmp2)
	//	y = 4
	//所以顺序就是：B D C A。
	//执行到 tmp1 时，输出：B 1 2 3。
	//执行到 tmp2 时，输出：D 3 2 5。
	//根据 defer 执行顺序原则，先声明的后执行，所以下一个该执行 C 了。
	//又因为传参是值赋值，所以在 A 的时候，无法用到 `x = 3` 和 `y = 4`，在 C 的时候，无法用到 `y = 4`。
	//执行到 C 时，输出：C 3 5 8
	//执行到 A 时，输出：A 1 3 4
	//到这，基本上 defer 就清楚了，大家可以根据自己的理解去记忆。
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, ret)
	return ret
}

func t1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

func t2() (a int) {
	defer func() {
		a++
	}()
	return 1
}

func t3() (b int) {
	a := 1
	defer func() {
		a++
	}()
	return 1
}
func t4() (a int) {
	defer func(a int) {
		a++ // 这个a++只会影响到传入的值，因为如果是传参他的值早就定下来了，为int的默认值0，所以加加之后还是1
	}(a)
	return 2
}

func goA() {
	defer (func() {
		if err := recover(); err != nil {
			fmt.Println("panic:" + fmt.Sprintf("%s", err))
		}
	})()
}

func goB() {
	panic("error")
}
