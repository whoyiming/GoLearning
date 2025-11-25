package main

import "fmt"

func main() {
	// 遍历切片
	slice := []string{"a", "b", "c"}
	for idx, val := range slice {
		fmt.Printf("切片索引：%d，值：%s\n", idx, val) // 0:a, 1:b, 2:c
	}

	// 遍历字符串（返回字节索引和Unicode码点）
	str := "Go语言"
	for idx, char := range str {
		fmt.Printf("字符串索引：%d，字符：%c（Unicode：%d）\n", idx, char, char)
		// 0:G(71), 1:o(111), 2:语(35821), 5:言(35328)（中文字符占3字节）
	}

	// 遍历map（返回键和值，无序）
	m := map[string]int{"a": 1, "b": 2}
	for key, val := range m {
		fmt.Printf("map键：%s，值：%d\n", key, val) // a:1, b:2（顺序不固定）
	}
}
