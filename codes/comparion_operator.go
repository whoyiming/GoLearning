package main

import "fmt"

func main() {
	// 数值比较
	m, n := 10, 20
	fmt.Println("m == n?", m == n) // false
	fmt.Println("m < n?", m < n)   // true
	fmt.Println("m >= n?", m >= n) // false

	// 字符串比较（按Unicode码点顺序）
	s1, s2 := "apple", "banana"
	s3 := "apple"
	fmt.Println("s1 == s3?", s1 == s3) // true
	fmt.Println("s1 < s2?", s1 < s2)   // true（'a' 的码点 < 'b'）

	// 布尔值比较
	b1, b2 := true, false
	fmt.Println("b1 != b2?", b1 != b2) // true

	// 指针比较（指向同一变量则相等）
	p1 := &m
	p2 := &m
	p3 := &n
	fmt.Println("p1 == p2?", p1 == p2) // true
	fmt.Println("p1 == p3?", p1 == p3) // false

	// 错误示例：不同类型不能直接比较
	// var num int = 10
	// var f float64 = 10.0
	// fmt.Println(num == f) // 编译错误：int 和 float64 不兼容
}
