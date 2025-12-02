package main

import "fmt"

// 场景1：引用传递 - 修改函数外部变量
func modifyValue(x int) {
    x = x * 2 // 只修改了函数内部的副本
}

func modifyValueByPointer(x *int) {
    *x = *x * 2 // 通过指针修改外部变量
}

// 场景2：大型结构体操作 - 避免拷贝开销
type LargeData struct {
    ID          int
    Name        string
    Data        [1000]int // 模拟大型数据
    Description string
}

// 值接收器 - 每次调用都会复制结构体
func (d LargeData) getSizeByValue() int {
    return len(d.Data)
}

// 指针接收器 - 不会复制结构体，只传递指针
func (d *LargeData) getSizeByPointer() int {
    return len(d.Data)
}

// 修改大型结构体内容
func updateLargeDataByPointer(d *LargeData, newValue int) {
    d.Data[0] = newValue // 直接修改原结构体
}

// 场景3：动态内存分配和可选返回值
func findElement(arr []int, target int) *int {
    for _, v := range arr {
        if v == target {
            return &v // 返回匹配元素的地址
        }
    }
    return nil // 未找到返回nil（空指针）
}

// 场景4：实现简单链表
type Node struct {
    Value int
    Next  *Node // 指向Next节点的指针
}

func createLinkedList(values []int) *Node {
    if len(values) == 0 {
        return nil
    }
    
    head := &Node{Value: values[0]} // 创建头节点
    current := head
    
    for i := 1; i < len(values); i++ {
        current.Next = &Node{Value: values[i]} // 链接新节点
        current = current.Next
    }
    
    return head
}

func printLinkedList(head *Node) {
    current := head
    fmt.Print("链表内容：")
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    fmt.Println("=== Go语言指针应用示例 ===")
    
    // 1. 引用传递示例
    fmt.Println("\n1. 引用传递示例:")
    num := 10
    fmt.Printf("修改前 num = %d\n", num)
    
    modifyValue(num)
    fmt.Printf("值传递后 num = %d（无变化）\n", num)
    
    modifyValueByPointer(&num)
    fmt.Printf("指针传递后 num = %d（被修改）\n", num)
    
    // 2. 大型结构体操作示例
    fmt.Println("\n2. 大型结构体操作示例:")
    largeData := LargeData{ID: 1, Name: "示例数据"}
    
    // 对比值接收器和指针接收器
    fmt.Printf("值接收器调用: %d\n", largeData.getSizeByValue())
    fmt.Printf("指针接收器调用: %d\n", (&largeData).getSizeByPointer())
    
    // Go语法糖：可以直接使用值类型调用指针接收器方法
    fmt.Printf("Go语法糖调用指针接收器: %d\n", largeData.getSizeByPointer())
    
    // 修改结构体内容
    updateLargeDataByPointer(&largeData, 999)
    fmt.Printf("修改后largeData.Data[0] = %d\n", largeData.Data[0])
    
    // 3. 动态内存分配和可选返回值示例
    fmt.Println("\n3. 动态内存分配和可选返回值示例:")
    arr := []int{1, 3, 5, 7, 9}
    
    // 查找存在的元素
    found := findElement(arr, 5)
    if found != nil {
        fmt.Printf("找到元素: %d\n", *found)
    } else {
        fmt.Println("未找到元素")
    }
    
    // 查找不存在的元素
    notFound := findElement(arr, 10)
    if notFound != nil {
        fmt.Printf("找到元素: %d\n", *notFound)
    } else {
        fmt.Println("未找到元素（返回nil）")
    }
    
    // 4. 链表实现示例
    fmt.Println("\n4. 链表实现示例:")
    values := []int{10, 20, 30, 40, 50}
    list := createLinkedList(values)
    printLinkedList(list)
}
