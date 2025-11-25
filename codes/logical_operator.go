package main

import "fmt"

func main() {
	// 基础逻辑运算
	a, b := true, false
	fmt.Println("a && b =", a && b) // false（且：都为true才true）
	fmt.Println("a || b =", a || b) // true（或：任一为true则true）
	fmt.Println("!a =", !a)         // false（非：取反）

	// 短路特性演示
	x, y := 10, 20

	// 逻辑与短路：左为false，右表达式（y++）不执行
	if x > 100 && y++ > 0 {
		fmt.Println("逻辑与执行")
	}
	fmt.Println("y =", y) // 20（y++未执行）

	// 逻辑或短路：左为true，右表达式（x++）不执行
	if x < 20 || x++ > 0 {
		fmt.Println("逻辑或执行")
	}
	fmt.Println("x =", x) // 10（x++未执行）
}

func autoincrement(){

}