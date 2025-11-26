package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 1. 基本的 switch
	fmt.Print("Go 运行在 ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// os 在这里仍然可见
		fmt.Printf("%s.\n", os)
	}

	// 2. 无表达式的 switch (相当于 if-else)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("上午好!")
	case t.Hour() < 17:
		fmt.Println("下午好!")
	default:
		fmt.Println("晚上好!")
	}

	// 3. case 可以有多个值
	day := "saturday"
	switch day {
	case "saturday", "sunday":
		fmt.Println("是周末！")
	default:
		fmt.Println("是工作日。")
	}

	// 4：fallthrough（穿透）
	num := 2
	switch num {
	case 2:
		fmt.Println("匹配2")
		fallthrough // 强制穿透到下一个case
	case 3:
		fmt.Println("穿透到3") // 会执行
	default:
		fmt.Println("默认")
	}
}
