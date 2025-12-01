package main

import (
	"errors"
	"fmt"
)

// 1. 基本函数：无参数、无返回值
func sayHello() {
	fmt.Println("=== 1. 基本函数 ===")
	fmt.Println("Hello, Go Functions!")
}

// 2. 带参数的函数：普通参数（值传递）
func add(a int, b int) int {
	// 同名参数可简写为 a, b int
	return a + b
}

// 3. 可变参数函数：参数数量不固定（本质是切片）
func sum(nums ...int) int {
	fmt.Println("=== 3. 可变参数函数 ===")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 4. 多返回值函数：Go 语言特色，可返回多个结果
func divide(a, b float64) (float64, error) {
	fmt.Println("=== 4. 多返回值函数 ===")
	if b == 0 {
		// 返回错误信息
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil // 正常返回结果和 nil 错误
}

// 5. 命名返回值函数：提前声明返回值变量，return 可省略参数
func calculate(a, b int) (sum int, product int) {
	fmt.Println("=== 5. 命名返回值函数 ===")
	sum = a + b // 直接给返回值变量赋值
	product = a * b
	return // 裸返回，自动返回 sum 和 product
}

// 6. 函数作为参数（高阶函数）
func processNumbers(nums []int, f func(int) int) []int {
	fmt.Println("=== 6. 函数作为参数 ===")
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = f(num) // 调用传入的函数
	}
	return result
}

// 7. 匿名函数：无函数名，直接定义并使用
func anonymousFuncDemo() {
	fmt.Println("=== 7. 匿名函数 ===")
	// 直接定义并调用
	func(msg string) {
		fmt.Println("匿名函数执行：", msg)
	}("Hello Anonymous!")

	// 赋值给变量后调用
	addFunc := func(a, b int) int {
		return a + b
	}
	fmt.Println("匿名函数变量调用：", addFunc(3, 4))
}

// 8. 闭包：函数嵌套，内部函数引用外部变量（延长外部变量生命周期）
func counter() func() int {
	fmt.Println("=== 8. 闭包 ===")
	count := 0 // 外部变量，被内部函数引用
	// 返回内部函数（闭包）
	return func() int {
		count++
		return count
	}
}

// 9. 递归函数：函数调用自身
func factorial(n int) int {
	fmt.Println("=== 9. 递归函数 ===")
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1) // 递归调用
}

// 10. defer 延迟执行：函数退出前执行（常用于资源释放）
func deferDemo() {
	fmt.Println("=== 10. defer 延迟执行 ===")
	defer fmt.Println("defer 1：最后执行") // 延迟到函数结束前
	defer fmt.Println("defer 2：倒数第二执行")
	fmt.Println("函数主体执行")
}

func main() {
	// 1. 基本函数调用
	sayHello()
	fmt.Println()

	// 2. 带参数函数调用
	fmt.Println("2. 3 + 5 =", add(3, 5))
	fmt.Println()

	// 3. 可变参数函数调用（可传任意个int参数）
	fmt.Println("3. 1+2+3+4 =", sum(1, 2, 3, 4))
	fmt.Println("3. 5+6 =", sum(5, 6))
	nums := []int{7, 8, 9}
	fmt.Println("3. 7+8+9 =", sum(nums...)) // 切片转可变参数
	fmt.Println()

	// 4. 多返回值函数调用（需要接收错误返回）
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("4. 错误：", err)
	} else {
		fmt.Println("4. 10 / 2 =", result)
	}
	// 测试错误场景
	_, err2 := divide(5, 0)
	if err2 != nil {
		fmt.Println("4. 错误场景：", err2)
	}
	fmt.Println()

	// 5. 命名返回值函数调用
	sumVal, productVal := calculate(4, 5)
	fmt.Println("5. 4+5 =", sumVal, "，4*5 =", productVal)
	fmt.Println()

	// 6. 高阶函数调用（传入自定义函数）
	numbers := []int{1, 2, 3, 4}
	// 传入匿名函数作为参数（求平方）
	squares := processNumbers(numbers, func(x int) int {
		return x * x
	})
	fmt.Println("6. 数组元素平方：", squares)
	// 传入匿名函数（求翻倍）
	doubles := processNumbers(numbers, func(x int) int {
		return x * 2
	})
	fmt.Println("6. 数组元素翻倍：", doubles)
	fmt.Println()

	// 7. 匿名函数演示
	anonymousFuncDemo()
	fmt.Println()

	// 8. 闭包演示（多次调用共享外部变量）
	c1 := counter()
	fmt.Println("8. 计数器1第一次：", c1())
	fmt.Println("8. 计数器1第二次：", c1())
	c2 := counter() // 新的闭包，独立变量
	fmt.Println("8. 计数器2第一次：", c2())
	fmt.Println()

	// 9. 递归函数演示（求阶乘）
	fmt.Println("9. 5的阶乘 =", factorial(5))
	fmt.Println()

	// 10. defer 演示
	deferDemo()
}
