package main

import "fmt"

// 泛型比较函数：返回两个值中较大的一个
// T：约束为 comparable（可比较类型）+ 支持 > 运算符（int/string/float64 等）
func Compare[T comparable](a, b T) bool {
	// 注意：comparable 仅保证支持 ==/!=，对于 >、< 运算符需类型本身支持（如数值/字符串）
	// 若传入不支持 > 的类型（如 struct），编译报错（类型安全）
	if a == b {
		return true
	}
	return false
}

func main() {
	fmt.Println("Compare(5, 5) =", Compare(5, 5))                               // 输出：true
	fmt.Println("Compare(3.14, 3.14) =", Compare(3.14, 3.14))                   // 输出：true
	fmt.Println("Compare(\"apple\", \"banana\") =", Compare("apple", "banana")) // 输出：false

}
