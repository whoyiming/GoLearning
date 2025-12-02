package main

import "fmt"

// 泛型求和函数
// T：类型参数，约束为 int | float64（仅允许这两种数值类型）
// nums：参数为 T 类型的切片
// 返回值：T 类型（与输入类型一致）
func Sum[T int | float64](nums []T) T {
	var total T // 声明 T 类型的零值（int=0，float64=0.0）
	for _, num := range nums {
		total += num // 仅数值类型支持 +=，约束保证了类型合法性
	}
	return total
}

func main() {
	// 1. 传入 int 类型切片（自动推导类型参数 T=int）
	intNums := []int{1, 2, 3, 4, 5}
	intSum := Sum(intNums)
	fmt.Printf("int 求和：%d\n", intSum) // 输出：15

	// 2. 传入 float64 类型切片（自动推导类型参数 T=float64）
	floatNums := []float64{1.1, 2.2, 3.3}
	floatSum := Sum(floatNums)
	fmt.Printf("float64 求和：%.2f\n", floatSum) // 输出：6.60

	// 3. 显式指定类型参数（推荐在类型不明确时使用）
	explicitSum := Sum[float64]([]float64{4.4, 5.5})
	fmt.Printf("显式指定类型求和：%.2f\n", explicitSum) // 输出：9.90
}
