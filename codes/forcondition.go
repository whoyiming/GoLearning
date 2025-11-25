package main

import "fmt"

func main() {
	// 1. 基本的 for 循环 (类似 C/Java)
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("0到9的和是:", sum)

	// 2. 省略初始化和后置语句 (类似 while)
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("大于100的最小2的幂是:", n)

	// 3. 无限循环 (需要 break 或 return 退出)
	// for {
	//     fmt.Println("这是一个无限循环，请手动停止程序")
	// }

	// 4. for range 遍历 (非常常用)
	s := "Go语言"
	for index, char := range s {
		// 注意 index 是字节位置，char 是 rune 类型
		fmt.Printf("在字节位置 %d 的字符是 %c\n", index, char)
	}
}
