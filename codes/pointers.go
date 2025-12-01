package main

import "fmt"

// 1. 基础指针：定义、取地址（&）、取值（*）
func basicPointer() {
	fmt.Println("=== 1. 基础指针操作 ===") // 输出：=== 1. 基础指针操作 ===
	// 普通变量
	a := 10
	// 输出示例：变量 a：值=10，地址=&0xc0000a6058（地址每次运行可能不同）
	fmt.Printf("变量 a：值=%d，地址=&%p\n", a, &a)

	// 声明指针变量（*int 表示 int 类型的指针）
	var p *int
	fmt.Printf("未赋值的指针 p：值=%v（nil）\n", p) // 输出：未赋值的指针 p：值=<nil>（nil）

	// 取变量 a 的地址，赋值给指针 p
	p = &a
	// 输出示例：指针 p：值=&0xc0000a6058（指向 a 的地址）
	fmt.Printf("指针 p：值=&%p（指向 a 的地址）\n", p)

	// 通过指针取值（*p 表示获取指针指向的变量的值）
	fmt.Printf("通过 *p 取值：%d（与 a 的值相同）\n", *p) // 输出：通过 *p 取值：10（与 a 的值相同）

	// 通过指针修改原变量的值
	*p = 20
	fmt.Printf("修改后：*p=%d，a=%d（原变量被修改）\n", *p, a) // 输出：修改后：*p=20，a=20（原变量被修改）
	fmt.Println()                                 // 输出：空行
}

// 2. 指针作为函数参数（实现“引用传递”，修改外部变量）
func modifyByPointer(x *int) {
	// x 是指针参数，接收外部变量的地址
	*x = *x * 2 // 通过指针修改外部变量的值
}

// 对比：值传递（无法修改外部变量）
func modifyByValue(x int) {
	x = x * 2 // 仅修改函数内部的副本
}

func pointerAsParam() {
	fmt.Println("=== 2. 指针作为函数参数 ===") // 输出：=== 2. 指针作为函数参数 ===
	b := 5
	fmt.Printf("修改前 b：%d\n", b) // 输出：修改前 b：5

	// 值传递：函数内部修改不影响外部
	modifyByValue(b)
	fmt.Printf("值传递后 b：%d（无变化）\n", b) // 输出：值传递后 b：5（无变化）

	// 指针传递：传递变量地址，函数内部可修改外部变量
	modifyByPointer(&b)
	fmt.Printf("指针传递后 b：%d（被修改）\n", b) // 输出：指针传递后 b：10（被修改）
	fmt.Println()                      // 输出：空行
}

// 3. 指针作为返回值（返回变量的地址）
// 注意：不能返回局部变量的指针（局部变量出栈后内存回收），这里返回全局变量/堆分配变量的指针
var globalVar = 100 // 全局变量（在全局区，不会随函数退出回收）

func getPointer() *int {
	// 返回全局变量的指针（安全）
	return &globalVar
}

// 或返回堆上分配的变量（通过 new 关键字，手动分配堆内存）
func newPointer() *int {
	// new(T) 分配 T 类型的内存，返回 *T 指针（堆内存，生命周期由 GC 管理）
	p := new(int)
	*p = 300
	return p
}

func pointerAsReturn() {
	fmt.Println("=== 3. 指针作为返回值 ===") // 输出：=== 3. 指针作为返回值 ===
	// 接收全局变量的指针
	p1 := getPointer()
	fmt.Printf("p1 指向全局变量：*p1=%d\n", *p1) // 输出：p1 指向全局变量：*p1=100

	// 接收堆上变量的指针
	p2 := newPointer()
	fmt.Printf("p2 指向堆变量：*p2=%d\n", *p2) // 输出：p2 指向堆变量：*p2=300

	// 修改返回的指针指向的值
	*p1 = 200
	*p2 = 400
	fmt.Printf("修改后：*p1=%d，globalVar=%d\n", *p1, globalVar) // 输出：修改后：*p1=200，globalVar=200
	fmt.Printf("修改后：*p2=%d\n", *p2)                         // 输出：修改后：*p2=400
	fmt.Println()                                           // 输出：空行
}

// 4. 结构体指针（常用场景：高效操作结构体，避免拷贝）
type Person struct {
	Name string
	Age  int
}

func structPointer() {
	fmt.Println("=== 4. 结构体指针 ===") // 输出：=== 4. 结构体指针 ===
	// 方式1：先定义结构体变量，再取地址
	person1 := Person{Name: "张三", Age: 25}
	p1 := &person1 // 结构体指针

	// 通过结构体指针访问字段（Go 语法糖：p1.Name 等价于 (*p1).Name）
	// 输出：p1 指向的结构体：Name=张三，Age=25
	fmt.Printf("p1 指向的结构体：Name=%s，Age=%d\n", p1.Name, (*p1).Age)

	// 通过结构体指针修改字段
	p1.Age = 26
	// 输出：修改后 person1：Name=张三，Age=26
	fmt.Printf("修改后 person1：Name=%s，Age=%d\n", person1.Name, person1.Age)

	// 方式2：直接创建结构体指针（使用 new）
	p2 := new(Person)
	p2.Name = "李四" // 自动解引用，无需 (*p2).Name
	p2.Age = 30
	// 输出：p2 指向的结构体：Name=李四，Age=30
	fmt.Printf("p2 指向的结构体：Name=%s，Age=%d\n", p2.Name, p2.Age)

	// 方式3：结构体字面量 + 取地址
	p3 := &Person{Name: "王五", Age: 35}
	// 输出：p3 指向的结构体：Name=王五，Age=35
	fmt.Printf("p3 指向的结构体：Name=%s，Age=%d\n", p3.Name, p3.Age)
	fmt.Println() // 输出：空行
}

// 5. 指针数组（数组元素是指针）
func pointerArray() {
	fmt.Println("=== 5. 指针数组 ===") // 输出：=== 5. 指针数组 ===
	// 普通数组（元素是值）
	arr1 := [3]int{1, 2, 3}
	fmt.Printf("普通数组 arr1：%v\n", arr1) // 输出：普通数组 arr1：[1 2 3]

	// 指针数组（元素是 *int 类型的指针）
	var arr2 [3]*int
	arr2[0] = &arr1[0] // 元素存储 arr1[0] 的地址
	arr2[1] = &arr1[1]
	arr2[2] = &arr1[2]

	// 输出示例：指针数组 arr2：[&0xc0000a60a8, &0xc0000a60b0, &0xc0000a60b8]
	fmt.Printf("指针数组 arr2：[&%p, &%p, &%p]\n", arr2[0], arr2[1], arr2[2])
	fmt.Printf("通过指针数组取值：[%d, %d, %d]\n", *arr2[0], *arr2[1], *arr2[2]) // 输出：通过指针数组取值：[1 2 3]

	// 通过指针数组修改原数组的值
	*arr2[0] = 10
	*arr2[1] = 20
	*arr2[2] = 30
	fmt.Printf("修改后 arr1：%v（原数组被修改）\n", arr1) // 输出：修改后 arr1：[10 20 30]（原数组被修改）
	fmt.Println()                             // 输出：空行
}

// 6. 数组指针（指向数组的指针）
func arrayPointer() {
	fmt.Println("=== 6. 数组指针 ===") // 输出：=== 6. 数组指针 ===
	arr := [3]int{10, 20, 30}
	// 声明数组指针（*[3]int 表示指向长度为 3 的 int 数组的指针）
	var p *[3]int = &arr

	// 输出示例：数组 arr：[10 20 30]，地址=&0xc0000a60c0
	fmt.Printf("数组 arr：%v，地址=&%p\n", arr, &arr)
	// 输出示例：数组指针 p：值=&0xc0000a60c0（指向 arr 的地址）
	fmt.Printf("数组指针 p：值=&%p（指向 arr 的地址）\n", p)

	// 通过数组指针访问元素（两种方式）
	fmt.Printf("p[0] = %d，(*p)[0] = %d\n", p[0], (*p)[0]) // 输出：p[0] = 10，(*p)[0] = 10

	// 通过数组指针修改数组元素
	p[1] = 200
	fmt.Printf("修改后 arr：%v\n", arr) // 输出：修改后 arr：[10 200 30]

	// 数组指针遍历
	fmt.Print("遍历数组指针 p：") // 输出：遍历数组指针 p：
	for i := 0; i < len(p); i++ {
		fmt.Printf("%d ", p[i]) // 输出：10 200 30
	}
	fmt.Println("\n") // 输出：换行 + 空行
}

// 7. 多级指针（指针的指针，较少用但需了解）
func multiLevelPointer() {
	fmt.Println("=== 7. 多级指针 ===") // 输出：=== 7. 多级指针 ===
	c := 100
	p1 := &c  // p1 是 *int（一级指针）
	p2 := &p1 // p2 是 **int（二级指针，指向一级指针 p1）
	p3 := &p2 // p3 是 ***int（三级指针，指向二级指针 p2）

	// 输出示例：c：值=100，地址=&0xc0000a6108
	fmt.Printf("c：值=%d，地址=&%p\n", c, &c)
	// 输出示例：p1（*int）：值=&0xc0000a6108，地址=&0xc0000c4018
	fmt.Printf("p1（*int）：值=&%p，地址=&%p\n", p1, &p1)
	// 输出示例：p2（**int）：值=&0xc0000c4018，地址=&0xc0000c4020
	fmt.Printf("p2（**int）：值=&%p，地址=&%p\n", p2, &p2)
	// 输出示例：p3（***int）：值=&0xc0000c4020
	fmt.Printf("p3（***int）：值=&%p\n", p3)

	// 多级指针取值（逐层解引用）
	fmt.Printf("*p1 = %d（一级解引用，取 c 的值）\n", *p1) // 输出：*p1 = 100（一级解引用，取 c 的值）
	fmt.Printf("**p2 = %d（二级解引用）\n", **p2)      // 输出：**p2 = 100（二级解引用）
	fmt.Printf("***p3 = %d（三级解引用）\n", ***p3)    // 输出：***p3 = 100（三级解引用）

	// 通过多级指针修改原变量
	***p3 = 200
	fmt.Printf("修改后 c = %d\n", c) // 输出：修改后 c = 200
	fmt.Println()                 // 输出：空行
}

func main() {
	// 按顺序执行所有指针示例
	basicPointer()      // 基础操作
	pointerAsParam()    // 指针作为参数
	pointerAsReturn()   // 指针作为返回值
	structPointer()     // 结构体指针
	pointerArray()      // 指针数组
	arrayPointer()      // 数组指针
	multiLevelPointer() // 多级指针
}
