package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"time"
)

// 函数定义
// - 函数用 `func` 声明。
// - 函数可以有一个或多个参数，需要有参数类型，用 `,` 分割。
// - 函数可以有一个或多个返回值，需要有返回值类型，用 `,` 分割。
// - 函数的参数是可选的，返回值也是可选的。
//
//	func function_name(input1 type1, input2 type2) (type1, type2) {
//	  // 函数体
//	  // 返回多个值
//	  return value1, value2
//	}
//
// ## 值传递
// 传递参数时，将参数复制一份传递到函数中，对参数进行调整后，不影响参数值。
// Go 语言默认是值传递。## 引用传递
//
//传递参数时，将参数的地址传递到函数中，对参数进行调整后，影响参数值。

func main() {
	str := "12345"
	fmt.Printf("MD5(%s): %s\n", str, MD5(str))
	fmt.Printf("current time str : %s\n", getTimeStr())
	fmt.Printf("current time unix : %d\n", getTimeInt())
	params := map[string]interface{}{
		"name": "Tom",
		"pwd":  "123456",
		"age":  30,
	}
	fmt.Printf("sign : %s\n", createSign(params))
}

// MD5方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 获取当前时间字符串
func getTimeStr() string {
	// 这里的数字对应Go的诞生时间（即2006年1月2日15点04分05秒MST），并且它们各自代表不同的时间组成部分
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前的时间戳
func getTimeInt() int64 {
	return time.Now().Unix()
}

// 生成签名
func createSign(params map[string]interface{}) string {
	var key []string
	var str = ""
	for k := range params {
		key = append(key, k)
	}
	sort.Strings(key)

	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&xl_%v=%v", key[i], params[key[i]])
		}
	}
	// 自定义密钥
	var secret string = "123456789"

	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(secret))
	return sign
}
