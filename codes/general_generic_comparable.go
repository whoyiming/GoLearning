package main

import "fmt"

// 泛型比较函数：返回两个值中较大的一个
// T：约束为 comparable（可比较类型）+ 支持 > 运算符（int/string/float64 等）
func Max[T comparable](a, b T) T {
	// 注意：comparable 仅保证支持 ==/!=，> 运算符需类型本身支持（如数值/字符串）
	// 若传入不支持 > 的类型（如 struct），编译报错（类型安全）
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("Max(3, 5) =", Max(3, 5))                               // 输出：5（T=int）
	fmt.Println("Max(3.14, 2.99) =", Max(3.14, 2.99))                   // 输出：3.14（T=float64）
	fmt.Println("Max(\"apple\", \"banana\") =", Max("apple", "banana")) // 输出：banana（T=string）

	// 编译错误：struct 虽为 comparable，但不支持 > 运算符
	// type Person struct{ Name string }
	// p1 := Person{Name: "张三"}
	// p2 := Person{Name: "李四"}
	// Max(p1, p2)
}
