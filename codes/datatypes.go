package main

import "fmt"

func main() {
	// 整数类型
	var a int = 10                                   // 通用整数（随系统）
	var b uint8 = 255                                // 无符号8位，最大值255（赋值256会编译错误）
	var c int64 = 9223372036854775807                // int64最大值
	fmt.Println("int:", a, "uint8:", b, "int64:", c) // 输出：int: 10 uint8: 255 int64: 9223372036854775807

	// 浮点类型
	var f1 float32 = 3.1415926                  // 32位浮点，精度有限
	f2 := 2.718281828459045                     // 短变量声明，默认float64
	fmt.Println("float32:", f1, "float64:", f2) // 输出：float32: 3.1415925 float64: 2.718281828459045

	// 注意：浮点数精度问题
	fmt.Println("0.1 + 0.2 =", 0.1+0.2) // 输出：0.1 + 0.2 = 0.30000000000000004（非0.3）

	// 复数类型
	var c1 complex64 = 3 + 4i                     // 32位复数（实部3，虚部4）
	c2 := 5 + 6i                                  // 短变量声明，默认complex128
	fmt.Println("complex64:", c1)                 // 输出：complex64: (3+4i)
	fmt.Println("实部:", real(c2), "虚部:", imag(c2)) // 输出：实部: 5 虚部: 6

	// 布尔类型
	var b1 bool = true // 显式声明
	b2 := false        // 短变量声明
	b3 := b1 && b2     // 逻辑与（都真才真）→ false
	b4 := b1 || b2     // 逻辑或（任一真则真）→ true
	b5 := !b1          // 逻辑非（取反）→ false

	fmt.Println("b1:", b1, "b2:", b2)                          // 输出：b1: true b2: false
	fmt.Println("b3 (&&):", b3, "b4 (||):", b4, "b5 (!):", b5) // 输出：b3 (&&): false b4 (||): true b5 (!): false

	// 条件判断（常用场景）
	if b1 {
		fmt.Println("b1 is true")
	} else {
		fmt.Println("b1 is false")
	}

	// 字符串类型
	var s1 string = "Hello, Go!" // 双引号定义
	s2 := "Hello\nWorld"         // 支持转义（\n换行）
	s3 := `line1
	line2
	line3` // 反引号多行字符串（不转义）
	s4 := s1 + " " + s2 // 字符串拼接（+号）

	fmt.Println("s1:", s1) // 输出：s1: Hello, Go!
	fmt.Println("s2:", s2) // 输出：s2: Hello
	//      World
	fmt.Println("s3:", s3) // 输出：s3: line1
	//      line2
	//      line3
	fmt.Println("拼接后s4:", s4) // 输出：拼接后s4: Hello, Go! Hello
	//      World
	fmt.Println("s1长度（字节数）:", len(s1))      // 输出：10（每个ASCII字符占1字节）
	fmt.Println("\"Go语言\"长度:", len("Go语言")) // 输出：8（Go占2字节，中文"语""言"各占3字节）

	// 特殊别名类型：byte 和 rune
	var ch1 byte = 'A'                                // ASCII字符，byte类型
	ch2 := '中'                                        // Unicode字符，自动推导为rune类型
	fmt.Println("byte('A'):", ch1, "rune('中'):", ch2) // 输出：byte('A'): 65 rune('中'): 20013（Unicode编码值）

	// 遍历字符串：byte遍历（中文乱码）vs rune遍历（正确显示）
	s := "Go语言"
	fmt.Println("\n按byte遍历（中文乱码）:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // 输出：G o è ¯ ­ è ¨ （中文按3字节拆分，乱码）
	}

	fmt.Println("\n按rune遍历（正确显示）:")
	for _, r := range s { // range会自动将字符串转为rune切片
		fmt.Printf("%c ", r) // 输出：G o 语 言
	}
}
