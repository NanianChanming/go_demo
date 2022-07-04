package main

import "fmt"

/*func main() {
	// 创建一个新的结构体
	//p1 := Person{"1", "张三", "1", 18}
	// 可以指定key -> value 的形式，忽略的字段为空
	//p2 := Person{id: "2", name: "李四", gender: "0"}
	//p2.age = 15
	//name := p1.name
	//fmt.Println(name)
	//fmt.Println(p1)
	//fmt.Println(p2)
	//fmt.Println(ParamStruct(p2))
	//fmt.Printf("p2 = %x", p2)
	//fmt.Println(p2)
	PointStruct()
}*/

// PointStruct 结构体指针
func PointStruct() {
	p := Person{id: "2", name: "李四", gender: "0"}
	//pp := &p
	printPerson(&p)
}

func printPerson(p *Person) {
	fmt.Println(p)
}

// ParamStruct 结构体作为函数参数传递
func ParamStruct(p Person) Person {
	p.name = "皇子"
	fmt.Printf("p = %x", p)
	return p
}

// Person go结构体
/*
go 结构体类似java中的类
定义语法：
结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
*/
type Person struct {
	id     string
	name   string
	gender string
	age    int
}
