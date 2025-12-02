package main

import (
	"errors"
	"fmt"
)

// 模拟除法运算，返回结果或错误
func divide(a, b int) (int, error) {
	if b == 0 {
		// 用 errors.New 创建基础错误
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil // 无错误时返回 nil
}

func main() {
	// 调用函数并检查错误
	result, err := divide(10, 2)
	if err != nil { // 错误不为 nil 表示执行失败
		fmt.Printf("除法失败：%v\n", err)
		return
	}
	fmt.Printf("10 / 2 = %d\n", result)

	// 测试错误场景
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("除法失败：%v\n", err2)
		return
	}
	fmt.Printf("10 / 0 = %d\n", result2) // 不会执行
}
