package main

import "fmt"

func main() {
	// 示例1：基础if-else
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 70 {
		fmt.Println("良好")
	} else {
		fmt.Println("及格/不及格")
	}

	// 示例2：if初始化语句（推荐用法）
	if age := 22; age >= 18 { // age仅在if-else内有效
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}
