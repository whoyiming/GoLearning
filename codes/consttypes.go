package main

import "fmt"

func main() {
	const LENGTH int = 20
	const WIDTH int = 15
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Println("面积为 : ", area)

	fmt.Println(a, b, c)
}
