package main

import "fmt"

func main() {
	// 1. 初始化 map（两种方式）
	m1 := make(map[string]int) // 空 map
	m2 := map[string]string{
		"name": "张三",
		"age":  "25",
		"city": "北京",
	}

	// 2. 新增/修改元素（key 存在则修改，不存在则新增）
	m1["math"] = 90
	m1["english"] = 85
	m2["age"] = "26"       // 修改已有 key
	fmt.Println("m1:", m1) // 输出：m1: map[english:85 math:90]
	fmt.Println("m2:", m2) // 输出：m2: map[age:26 city:北京 name:张三]

	// 3. 查询元素（第二个返回值表示 key 是否存在）
	score, ok := m1["math"]
	if ok {
		fmt.Println("math 成绩:", score) // 输出：math 成绩: 90
	} else {
		fmt.Println("math 成绩不存在")
	}

	// 4. 删除元素（delete(map, key)，key 不存在则无操作）
	delete(m2, "city")
	fmt.Println("m2 删除 city 后:", m2) // 输出：m2 删除 city 后: map[age:26 name:张三]

	// 5. 遍历 map（无序，每次遍历顺序可能不同）
	fmt.Println("遍历 m1:")
	for k, v := range m1 {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// 6. map 是引用类型（拷贝后修改会影响原 map）
	m3 := m1
	m3["english"] = 95
	fmt.Println("原 m1:", m1)   // 输出：原 m1: map[english:95 math:90]
	fmt.Println("拷贝后 m3:", m3) // 输出：拷贝后 m3: map[english:95 math:90]
}
