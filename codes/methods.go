package main

import "fmt"

// 1. 定义结构体（方法最常用的绑定类型）
type Rectangle struct {
	Width  float64
	Height float64
}

// 2. 结构体 + 值接收者方法（不能修改原结构体的值）
// 计算面积：值接收者接收的是结构体副本
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 3. 结构体 + 指针接收者方法（可以修改原结构体的值）
// 缩放矩形：指针接收者接收的是结构体地址，直接操作原数据
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// 4. 自定义类型（类型别名）+ 方法
// 给 int 起别名（基础类型不能直接绑定方法，需通过类型别名）
type MyInt int

// 自定义类型的方法（值接收者）
func (m MyInt) Add(n MyInt) MyInt {
	return m + n
}

// 自定义类型的指针接收者方法
func (m *MyInt) Double() {
	*m *= 2
}

func main() {
	// 示例1：结构体方法调用
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("=== 结构体方法 ===")
	fmt.Printf("原始矩形：宽=%.1f，高=%.1f\n", rect.Width, rect.Height)

	// 调用值接收者方法（无修改）
	area := rect.Area()
	fmt.Printf("矩形面积：%.1f\n", area) // 输出：15.0

	// 调用指针接收者方法（修改原结构体）
	rect.Scale(2)                                                // Go语法糖：rect是值类型，调用指针方法时自动转&rect
	fmt.Printf("缩放2倍后：宽=%.1f，高=%.1f\n", rect.Width, rect.Height) // 输出：10.0, 6.0
	fmt.Printf("缩放后面积：%.1f\n", rect.Area())                      // 输出：60.0
	fmt.Println()

	// 示例2：自定义类型方法调用
	fmt.Println("=== 自定义类型方法 ===")
	var a MyInt = 10
	var b MyInt = 20

	// 调用值接收者方法（无修改）
	sum := a.Add(b)
	fmt.Printf("a + b = %d\n", sum) // 输出：30

	// 调用指针接收者方法（修改原变量）
	a.Double()                  // 语法糖：a是值类型，自动转&a
	fmt.Printf("a 翻倍后：%d\n", a) // 输出：20

	// 直接用指针调用（等价于上面的语法糖）
	var c MyInt = 5
	(&c).Double()
	fmt.Printf("c 翻倍后：%d\n", c) // 输出：10
}
