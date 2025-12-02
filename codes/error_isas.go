package main

import (
	"errors"
	"fmt"
)

// 定义自定义错误类型（实现 error 接口）
type NotFoundError struct {
	Resource string // 资源名称（如用户、文件）
}

// 实现 Error() 方法，满足 error 接口
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s 不存在", e.Resource)
}

// 模拟查询资源
func findResource(name string) error {
	if name == "" {
		return errors.New("资源名称不能为空")
	}
	// 返回自定义错误类型
	return fmt.Errorf("查询失败：%w", &NotFoundError{Resource: name})
}

func main() {
	// 场景1：用 errors.Is 判断基础错误
	err1 := findResource("")
	if errors.Is(err1, errors.New("资源名称不能为空")) {
		fmt.Printf("错误匹配：%v\n", err1)
	}

	// 场景2：用 errors.As 转换自定义错误类型
	err2 := findResource("订单10086")
	var notFoundErr *NotFoundError
	if errors.As(err2, &notFoundErr) { // 检查错误链中是否有 NotFoundError 类型
		fmt.Printf("自定义错误：资源类型=%s，错误信息=%v\n", notFoundErr.Resource, err2)
	}

	// 场景3：嵌套包装的错误判断
	err3 := fmt.Errorf("处理订单失败：%w", err2)
	if errors.As(err3, &notFoundErr) {
		fmt.Printf("嵌套错误匹配：%v\n", err3)
	}
}
