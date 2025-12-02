package main

import "fmt"

// 泛型结构体：Stack（栈），T 为元素类型
type Stack[T any] struct {
	elements []T
}

// 泛型方法：Push（入栈）- 使用指针接收器
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// 泛型方法：Len（获取栈长度）- 使用值接收器
func (s Stack[T]) Len() int {
	return len(s.elements)
}

func main() {
	// 示例1：使用指针创建栈实例
	fmt.Println("=== 使用指针创建栈实例 (&Stack[string]{}) ===")
	strStack1 := &Stack[string]{}
	strStack1.Push("a")
	fmt.Printf("栈1长度: %d\n", strStack1.Len())
	
	// 复制指针变量（两者指向同一底层数据）
	strStackPtrCopy := strStack1
	strStackPtrCopy.Push("b")
	fmt.Printf("复制指针后，栈1长度: %d\n", strStack1.Len())
	
	// 示例2：不使用指针创建栈实例
	fmt.Println("\n=== 不使用指针创建栈实例 (Stack[string]{}) ===")
	strStack2 := Stack[string]{}
	strStack2.Push("x")
	fmt.Printf("栈2长度: %d\n", strStack2.Len())
	
	// 复制值变量（创建独立副本）
	strStackValCopy := strStack2
	strStackValCopy.Push("y")
	fmt.Printf("复制值后，原栈2长度: %d\n", strStack2.Len())
	fmt.Printf("复制后新栈长度: %d\n", strStackValCopy.Len())
	
	// 示例3：方法接收器类型的影响
	fmt.Println("\n=== 方法接收器类型的影响 ===")
	// 无论栈是指针还是值，都可以调用指针接收器方法（Go会自动转换）
	var stack Stack[int]
	stack.Push(1) // 自动转换为 (&stack).Push(1)
	fmt.Printf("使用值调用指针接收器方法后的长度: %d\n", stack.Len())
}