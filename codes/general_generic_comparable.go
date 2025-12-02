package main

import (
	"constraints"
	"fmt"
)

// 泛型比较函数：返回两个值中较大的一个
// T：约束为 constraints.Ordered（支持 > 运算符的有序类型，如 int/float64/string 等）
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Max(3, 5) =", Max(3, 5))                               // 输出：5（T=int）
	fmt.Println("Max(3.14, 2.99) =", Max(3.14, 2.99))                   // 输出：3.14（T=float64）
	fmt.Println("Max(\"apple\", \"banana\") =", Max("apple", "banana")) // 输出：banana（T=string）

	// 编译错误（保留原注释，验证类型安全）：struct 不属于 constraints.Ordered，无法传入
	// type Person struct{ Name string }
	// p1 := Person{Name: "张三"}
	// p2 := Person{Name: "李四"}
	// Max(p1, p2)
}
