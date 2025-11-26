package main

import "fmt"

// 1. 定义结构体（首字母大写表示可导出，供其他包使用，字段首字母大写表示可导出，供其他包使用，字段首字母小写表示包内私有，不导出）
type Person struct {
	Name string
	Age  int
	Addr string
}

// 2. 定义结构体方法（值接收者：修改不会影响原结构体）
func (p Person) SayHello() {
	fmt.Printf("大家好，我是 %s，今年 %d 岁\n", p.Name, p.Age)
}

// 3. 指针接收者：修改会影响原结构体（推荐用于大结构体，避免拷贝）
func (p *Person) GrowUp() {
	p.Age++ // 等价于 (*p).Age++
}

func main() {
	// 4. 创建结构体实例（三种方式）
	p1 := Person{"张三", 25, "北京"}      // 按字段顺序初始化
	p2 := Person{Name: "李四", Age: 30} // 指定字段初始化（未指定字段为零值）
	p3 := new(Person)                 // new 返回指针 (*Person)
	p3.Name = "王五"
	p3.Age = 28

	fmt.Println("p1:", p1) // 输出：p1: {张三 25 北京}
	fmt.Println("p2:", p2) // 输出：p2: {李四 30 }（Addr 为零值 ""）
	fmt.Println("p3:", p3) // 输出：p3: &{王五 28 }（指针类型）

	// 5. 访问字段（指针结构体可直接通过 . 访问，无需解引用）
	fmt.Println("p3 姓名:", p3.Name) // 输出：p3 姓名: 王五

	// 6. 调用方法（值接收者和指针接收者均可直接调用，Go 自动转换）
	p1.SayHello() // 输出：大家好，我是 张三，今年 25 岁
	p3.GrowUp()   // 调用指针方法，修改原结构体年龄
	p3.SayHello() // 输出：大家好，我是 王五，今年 29 岁

	// 7. 结构体嵌套（实现“组合”，类似继承）
	type Student struct {
		Person // 嵌套结构体（匿名字段，继承 Person 的字段和方法）
		Score  int
	}
	stu := Student{Person: Person{Name: "赵六", Age: 20}, Score: 95}
	fmt.Println("学生姓名:", stu.Name) // 直接访问嵌套字段
	stu.SayHello()                 // 调用嵌套结构体的方法
}
