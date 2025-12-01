package main

import "fmt"

// 1. 定义接口：Shape（形状），声明方法签名
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// 2. 实现类1：Circle（圆形），隐式实现Shape接口
type Circle struct {
	Radius float64
}

// 实现Shape的Area方法
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 实现Shape的Perimeter方法（必须实现接口的所有方法）
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 3. 实现类2：Rectangle（矩形），隐式实现Shape接口
type Rectangle struct {
	Width  float64
	Height float64
}

// 实现Shape的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 实现Shape的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 4. 接口的多态应用：接收Shape接口参数，自动适配不同实现类
func PrintShapeInfo(s Shape) {
	fmt.Printf("面积：%.2f，周长：%.2f\n", s.Area(), s.Perimeter())
}

// 5. 空接口：interface{}（无任何方法），可接收任意类型
func PrintAnyValue(v interface{}) {
	fmt.Printf("类型：%T，值：%v\n", v, v)
}

// 6. 接口断言：判断接口的具体实现类型
func GetShapeType(s Shape) string {
	// 格式：value, ok := 接口变量.(具体类型)
	if _, ok := s.(Circle); ok {
		return "圆形"
	}
	if _, ok := s.(Rectangle); ok {
		return "矩形"
	}
	return "未知形状"
}

func main() {
	// 示例1：多态调用（同一接口接收不同类型）
	fmt.Println("=== 接口多态 ===")
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	fmt.Print("圆形：")
	PrintShapeInfo(circle) // 输出：面积：78.50，周长：31.40
	fmt.Print("矩形：")
	PrintShapeInfo(rectangle) // 输出：面积：12.00，周长：14.00
	fmt.Println()

	// 示例2：接口断言（判断具体类型）
	fmt.Println("=== 接口断言 ===")
	fmt.Printf("circle 是 %s\n", GetShapeType(circle))       // 输出：圆形
	fmt.Printf("rectangle 是 %s\n", GetShapeType(rectangle)) // 输出：矩形
	fmt.Println()

	// 示例3：空接口（接收任意类型）
	fmt.Println("=== 空接口 ===")
	PrintAnyValue(100)        // 输出：类型：int，值：100
	PrintAnyValue("Hello Go") // 输出：类型：string，值：Hello Go
	PrintAnyValue(true)       // 输出：类型：bool，值：true
	PrintAnyValue(circle)     // 输出：类型：main.Circle，值：{5}
}
