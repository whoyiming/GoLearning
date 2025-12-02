package main

import "fmt"

// 泛型 Map：K 为键类型（需 comparable），V 为值类型（any）
type MyMap[K comparable, V any] struct {
	data map[K]V
}

// 初始化泛型 Map
func NewMyMap[K comparable, V any]() *MyMap[K, V] {
	return &MyMap[K, V]{
		data: make(map[K]V),
	}
}

// 存值
func (m *MyMap[K, V]) Set(key K, val V) {
	m.data[key] = val
}

// 取值
func (m *MyMap[K, V]) Get(key K) (V, bool) {
	val, ok := m.data[key]
	return val, ok
}

func main() {
	// 键为 string，值为 int 的 Map
	strIntMap := NewMyMap[string, int]()
	strIntMap.Set("a", 1)
	strIntMap.Set("b", 2)
	val1, ok1 := strIntMap.Get("a")
	fmt.Printf("key=a：值=%d，存在？%t\n", val1, ok1) // 输出：1，true

	// 键为 int，值为 Person 的 Map
	type Person struct{ Age int }
	intPersonMap := NewMyMap[int, Person]()
	intPersonMap.Set(1001, Person{Age: 25})
	val2, ok2 := intPersonMap.Get(1001)
	fmt.Printf("key=1001：值=%+v，存在？%t\n", val2, ok2) // 输出：{Age:25}，true
}
