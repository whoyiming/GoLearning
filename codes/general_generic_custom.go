package main

import "fmt"

// 自定义约束：仅允许数值类型（int/int64/float32/float64）
type Number interface {
	int | int64 | float32 | float64
}

// 泛型函数：计算切片平均值（仅支持 Number 约束的类型）
func Average[T Number](nums []T) float64 {
	var total T
	for _, num := range nums {
		total += num
	}
	// 转换为 float64 返回平均值
	return float64(total) / float64(len(nums))
}

func main() {
	int64Nums := []int64{10, 20, 30, 40}
	fmt.Printf("int64 切片平均值：%.2f\n", Average(int64Nums)) // 输出：25.00

	float32Nums := []float32{1.5, 2.5, 3.5}
	fmt.Printf("float32 切片平均值：%.2f\n", Average(float32Nums)) // 输出：2.50

	// 编译错误：string 不满足 Number 约束
	// strNums := []string{"a", "b"}
	// Average(strNums)
}
