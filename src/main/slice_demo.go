package main

import "fmt"

/*func main() {
	// LenFunc()
	// NilSlice()
	// SplitSlice()
	AppendAndCopy()
}*/

func AppendAndCopy() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	printSliceInt(slice)
	slice = append(slice, 7)
	printSliceInt(slice)

	// 创建一个长度及容量一样的切片
	var s1 = make([]int, len(slice), cap(slice))
	// 执行函数copy
	copy(s1, slice)
	printSliceInt(s1)

}

func SplitSlice() {
	slice := make([]string, 5, 5)
	slice[0] = "德"
	slice[1] = "玛"
	slice[2] = "西"
	slice[3] = "亚"
	slice[4] = "皇子"
	printSlice(slice)

	printSlice(slice[0:3])
	printSlice(slice[3:])
}

func NilSlice() {
	var ns []string
	printSlice(ns)
	if ns == nil {
		fmt.Println("空切片")
	}
}

func LenFunc() {
	slice := make([]string, 5, 5)
	slice[0] = "hello"
	// slice = append(slice, "hello", "world")
	printSlice(slice)
}

func SliceInit() {
	// 切片初始化
	// 直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
	// s := []int{1, 2, 3}

	// arr := []string{"hello", "world"}
	// 初始化切片 s，是数组 arr 的引用。
	// s := arr[:]

	// 将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。
	// s := arr[startIndex:endIndex]

	// 默认 endIndex 时将表示一直到arr的最后一个元素。
	// s := arr[startIndex:]

	// 默认 startIndex 时将表示从 arr 的第一个元素开始。
	// s := arr[:endIndex]

	// 通过切片 s 初始化切片 s1。
	// s1 := s[startIndex:endIndex]

	// 通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。
	// s := make([]int,len,cap)
}

// Go 语言切片是对数组的抽象。
// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
// Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
// 类似java中list

func SliceDemo() {
	// 声明一个未定义长度的数组来定义切片 var identifier []type
	// 切片不需要说明长度。
	// var slice []string
	// 或使用 make() 函数来创建切片:
	// var slice []string = make([]string, 10)
	// 也可以简写成
	slice := make([]string, 10)
	fmt.Println(slice)

	// 也可以指定容量，其中 capacity 为可选参数。
	// make([]T, length, capacity)
}

func printSlice(x []string) {
	fmt.Printf("slice = %v, len = %d, cap = %d \n", x, len(x), cap(x))
}

func printSliceInt(x []int) {
	fmt.Printf("slice = %v, len = %d, cap = %d \n", x, len(x), cap(x))
}
