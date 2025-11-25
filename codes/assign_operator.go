package main

import "fmt"

func main() {
	// 基础赋值
	var x int = 10
	x = 20
	fmt.Println("x =", x) // 20

	// 复合赋值
	y := 10
	y += 5                // y = 10 + 5 = 15
	y *= 2                // y = 15 * 2 = 30
	y %= 7                // y = 30 % 7 = 2
	fmt.Println("y =", y) // 2

	// 位运算复合赋值
	z := 0b1010                                   // 10（二进制）
	z <<= 2                                       // 左移2位 → 0b101000 = 40
	fmt.Printf("z <<= 2 → %d（二进制：%b）\n", z, z) // 40（101000）
}
