package main

import "fmt"

/*func main() {
	//i := 1
	//fmt.Printf("%x", &i)
	//point1()
	//NullPoint()
	ArrPoint()
}*/

func point1() {
	// 声明指针 var var_name *var-type
	// 声明指向int指针
	var p *int
	/*
		指针使用:
			定义指针变量。
			为指针变量赋值。
			访问指针变量中指向地址的值。
	*/
	i := 1
	p = &i
	fmt.Printf("a 变量的地址是: %x\n", &i)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", p)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *p)

	a := i + *p
	fmt.Println(a)
}

func NullPoint() {
	var np *string
	fmt.Printf("%x \n", np)

	if np == nil {
		fmt.Println(true)
	}

	if np != nil {
		fmt.Println(false)
	}
}

func ArrPoint() {
	arr := [3]string{"hello", "world", "go"}
	var ap [3]*string

	for i := 0; i < len(arr); i++ {
		ap[i] = &arr[i]
	}

	// s指向同一个变量地址，指针指向的地址为同一个
	/*for i, _ := range arr {
		s := arr[i]
		ap[i] = &s
	}*/

	fmt.Println("遍历指针数组:------")
	for _, s := range ap {
		fmt.Println(*s)
	}
}
