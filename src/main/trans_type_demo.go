package main

import "fmt"

/*func main() {
	transType(10)
}*/

// 类型转换
// type_name(value)
func transType(x int) {
	fmt.Printf("原始类型为：%T\n", x)
	s := string(x)
	fmt.Printf("转换后类型为：%T\n", s)
}
