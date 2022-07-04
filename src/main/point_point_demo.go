package main

import "fmt"

/*func main() {
	//PointPoint()
	var ps *string
	var pi *int
	a, b := PointParam(ps, pi)
	fmt.Printf("%s%d", *a, *b)
}*/

// PointParam /* 指针可以当作参数传递，只需要在函数中声明指针类型参数即可
func PointParam(x *string, y *int) (*string, *int) {
	s, i := "hello g", 0
	x, y = &s, &i
	return x, y
}

func PointPoint() {
	// 指向指针的指针
	// 声明格式: var pp **type
	str := "hello world"
	var ps *string
	var ptp **string
	ps = &str

	ptp = &ps
	fmt.Printf("变量地址 = %x 值 = %s \n", str, str)
	fmt.Printf("指针地址 = %x 值 = %s \n", *ps, **ptp)
	fmt.Printf("指向指针的指针地址 = %x 值 = %s \n", **ptp, **ptp)
}
