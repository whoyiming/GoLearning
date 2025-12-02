package main

import "fmt"

// 泛型结构体：Stack（栈），T 为元素类型（any 支持所有类型）
type Stack[T any] struct {
	elements []T // 存储 T 类型的元素
}

// 泛型方法：Push（入栈）
func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

// 泛型方法：Pop（出栈），返回元素和是否成功
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // 返回 T 类型的零值
		return zero, false
	}
	// 取最后一个元素
	lastIdx := len(s.elements) - 1
	item := s.elements[lastIdx]
	s.elements = s.elements[:lastIdx]
	return item, true
}

// 泛型方法：Len（获取栈长度）
func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func main() {
	// 1. 字符串类型栈（T=string）
	strStack := &Stack[string]{}
	strStack.Push("a")
	strStack.Push("b")
	strStack.Push("c")
	fmt.Printf("字符串栈长度：%d\n", strStack.Len()) // 输出：3

	item1, ok1 := strStack.Pop()
	fmt.Printf("出栈元素：%s，成功？%t\n", item1, ok1) // 输出：c，true

	// 2. int 类型栈（T=int）
	intStack := &Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	item2, ok2 := intStack.Pop()
	fmt.Printf("出栈元素：%d，成功？%t\n", item2, ok2) // 输出：20，true

	// 3. 自定义类型栈（T=Person）
	type Person struct{ Name string }
	personStack := &Stack[Person]{}
	personStack.Push(Person{Name: "张三"})
	item3, ok3 := personStack.Pop()
	fmt.Printf("出栈元素：%+v，成功？%t\n", item3, ok3) // 输出：{Name:张三}，true
}
