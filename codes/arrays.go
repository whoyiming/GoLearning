package main

import "fmt"

func main() {
	//定义方式
	// 1. 声明并初始化（指定长度）
	var arr1 [3]int = [3]int{1, 2, 3}
	// 2. 短变量声明（自动推导类型）
	arr2 := [3]int{4, 5, 6}
	// 3. 自动推断长度（... 表示由初始化值决定）
	arr3 := [...]int{7, 8, 9}
	// 4. 部分初始化（未指定的元素为类型零值）
	arr4 := [5]string{"a", "b", 3: "d"} // 索引 0:a,1:b,2:"",3:d,4:""
	//基本操作
	arr := [3]int{10, 20, 30}

	// 1. 索引访问（0 开始）
	fmt.Println(arr[0]) // 输出 10
	arr[1] = 200        // 修改元素

	// 2. 遍历（for 循环）
	for i := 0; i < len(arr); i++ {
		fmt.Printf("索引 %d: %d\n", i, arr[i])
	}

	// 3. 遍历（for range，推荐）
	for idx, val := range arr {
		fmt.Printf("idx: %d, val: %d\n", idx, val)
	}
}
