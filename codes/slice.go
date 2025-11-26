package main

import "fmt"

func main() {
	// 1. 声明切片（三种方式）
	var s1 []int                  // 空切片（len=0, cap=0）
	s2 := []string{"a", "b", "c"} // 直接初始化
	s3 := make([]float64, 3, 5)   // make(类型, 长度, 容量)，未初始化元素为零值

	fmt.Println("s1:", s1, "len:", len(s1), "cap:", cap(s1)) // s1: [] len: 0 cap: 0
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2)) // s2: [a b c] len: 3 cap: 3
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3)) // s3: [0 0 0] len: 3 cap: 5

	// 2. 切片截取（左闭右开，不修改原切片，共享底层数组）
	s4 := s2[1:3]              // 从 s2 截取索引 1~2（不含3）
	fmt.Println("s4:", s4)     // 输出：s4: [b c]
	s4[0] = "x"                // 修改截取切片，原切片也会变（共享底层数组）
	fmt.Println("修改后 s2:", s2) // 输出：修改后 s2: [a x c]

	// 3. append 扩容（容量不足时自动扩容，生成新底层数组）
	s5 := []int{1, 2}
	fmt.Println("s5 初始:", s5, "len:", len(s5), "cap:", cap(s5))  // [1 2] len:2 cap:2
	s5 = append(s5, 3, 4, 5)                                     // 追加元素，容量不足时扩容（通常翻倍）
	fmt.Println("s5 追加后:", s5, "len:", len(s5), "cap:", cap(s5)) // [1 2 3 4 5] len:5 cap:6

	// 4. 切片拷贝（copy(dst, src)，仅拷贝元素，不共享底层数组）
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 999
	fmt.Println("原 src:", src)   // 输出：原 src: [10 20 30]
	fmt.Println("拷贝后 dst:", dst) // 输出：拷贝后 dst: [999 20 30]
}
