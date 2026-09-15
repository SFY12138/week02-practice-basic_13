package main

import "fmt"

// Divide 函数：执行除法操作
func Divide(a, b float64) float64 {
	if b == 0 {
		panic("除数不能为零")
	}
	return a / b
}

// SafeDivide 函数：使用 recover 捕获 panic
func SafeDivide(a, b float64) (float64, string) {
	var result float64
	var errMsg string

	func() {
		defer func() {
			if r := recover(); r != nil {
				errMsg = fmt.Sprintf("错误: %v", r)
			}
		}()
		result = Divide(a, b)
	}()

	return result, errMsg
}

func main() {
	// 测试正常情况
	result, err := SafeDivide(10.0, 2.0)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("10.0 / 2.0 =", result)
	}

	// 测试除数为零的情况
	result, err = SafeDivide(10.0, 0.0)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Println("10.0 / 0.0 =", result)
	}
}
