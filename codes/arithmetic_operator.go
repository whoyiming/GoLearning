package main

import "fmt"

func main() {
	// 基础算术运算
	a, b := 10, 3
	fmt.Println("a + b =", a+b)   // 13
	fmt.Println("a - b =", a-b)   // 7
	fmt.Println("a * b =", a*b)   // 30
	fmt.Println("a / b =", a/b)   // 3（整数除法舍弃小数）
	fmt.Println("a % b =", a%b)   // 1（余数）
	fmt.Println("-a % b =", -a%b) // -1（符号与被除数一致）

	// 浮点数运算
	f1, f2 := 10.0, 3.0
	fmt.Println("f1 / f2 =", f1/f2) // 3.3333333333333335

	// 字符串拼接（仅 + 支持）
	s1, s2 := "Go", "Lang"
	fmt.Println("s1 + s2 =", s1+s2) // GoLang

	// 自增自减（只能独立语句，不能赋值）
	i, j := 5, 5
	i++ // 正确：i = 6
	j-- // 正确：j = 4
	// fmt.Println(i++)  // 错误：不能作为表达式使用
	// k := i++          // 错误：不能赋值
	fmt.Println("i =", i, "j =", j) // 6 4
}
