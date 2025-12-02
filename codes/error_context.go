package main

import (
	"errors"
	"fmt"
)

// 模拟查询用户
func getUserID(username string) (int, error) {
	if username == "" {
		return 0, errors.New("用户名不能为空")
	}
	if username != "admin" {
		return 0, errors.New("用户不存在")
	}
	return 1001, nil
}

// 模拟查询用户信息（包装下层错误）
func getUserInfo(username string) (string, error) {
	userID, err := getUserID(username)
	if err != nil {
		// 用 %w 包装原始错误，添加上下文
		return "", fmt.Errorf("查询用户ID失败：%w", err)
	}
	return fmt.Sprintf("用户ID：%d，用户名：%s", userID, username), nil
}

func main() {
	info, err := getUserInfo("guest")
	if err != nil {
		fmt.Printf("获取用户信息失败：%v\n", err)

		// 解包原始错误并判断
		originalErr := errors.Unwrap(err)
		fmt.Printf("原始错误：%v\n", originalErr)
		return
	}
	fmt.Printf("用户信息：%s\n", info)
}
