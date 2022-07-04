package main

import "fmt"

/*func main() {
	//arr()
	Twoarr()
}*/

func arr() {
	/*
		数组定义：var name = [长度]type{值}
		不确定长度可以用...代替，会自动根据值个数推导长度
		初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	*/
	var arr = [...]string{"hello", "go", "world"}
	// 遍历数组
	for i, s := range arr {
		fmt.Printf("arr[%d] = %s \n", i, s)
	}
	fmt.Println("长度 = ", len(arr))

	// 可以指定索引位置的值
	balance := [5]float32{1: 2.0, 3: 7.0}
	balance[4] = 30
	for i, s := range balance {
		fmt.Printf("balance[%d] = %f \n", i, s)
	}
}

func Twoarr() {
	// 多维数组
	// var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
	var arrs [2][3]string
	arrs[0][0] = "hello"
	arrs[0][1] = "go"
	arrs[0][2] = "world"

	for i, a := range arrs {
		for j, s := range a {
			fmt.Printf("arrs[%d][%d] = %s \n", i, j, s)
		}
	}

	//多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组
	a := [2][3]string{
		{1: "hello", "go"},
		{"德玛", "西亚", "皇子"},
	}

	for i, s := range a[0] {
		fmt.Printf("a[0][%d] = %s \n", i, s)
	}

	for i, s := range a[1] {
		fmt.Printf("a[1][%d] = %s \n", i, s)
	}

}
