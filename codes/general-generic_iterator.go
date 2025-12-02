package main

import "fmt"

// 泛型接口：Iterator（迭代器），T 为迭代元素类型
type Iterator[T any] interface {
	Next() (T, bool) // 返回下一个元素和是否迭代完成
}

// 泛型切片迭代器（实现 Iterator 接口）
type SliceIterator[T any] struct {
	slice []T
	index int
}

// 实现 Iterator 接口的 Next 方法
func (s *SliceIterator[T]) Next() (T, bool) {
	if s.index >= len(s.slice) {
		var zero T
		return zero, false
	}
	val := s.slice[s.index]
	s.index++
	return val, true
}

// 工厂函数：创建切片迭代器
func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice: slice,
		index: 0,
	}
}

// 通用迭代函数：接收任意类型的 Iterator
func Iterate[T any](it Iterator[T]) {
	fmt.Print("迭代结果：")
	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func main() {
	// int 切片迭代
	intSlice := []int{1, 2, 3, 4}
	intIt := NewSliceIterator(intSlice)
	Iterate(intIt) // 输出：迭代结果：1 2 3 4

	// string 切片迭代
	strSlice := []string{"apple", "banana", "cherry"}
	strIt := NewSliceIterator(strSlice)
	Iterate(strIt) // 输出：迭代结果：apple banana cherry
}
