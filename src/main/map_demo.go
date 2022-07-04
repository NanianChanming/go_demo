package main

import "fmt"

/*func main() {
	defMap()
}*/

/*
Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
*/
func defMap() {
	// 可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
	/* 声明变量，默认 map 是 nil */
	// var map_variable map[key_data_type]value_data_type
	var m map[string]string
	if m == nil {
		fmt.Printf("空map \n")
	}

	/* 使用 make 函数 */
	// map_variable := make(map[key_data_type]value_data_type)
	m1 := make(map[string]string)
	m1["dema"] = "xiya"
	m1["nuoke"] = "sasi"
	s, ok := m1["dema"]
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("key 不存在")
	}

	s1, ok := m1["zuan"]
	if ok {
		fmt.Println(s1)
	} else {
		fmt.Println("key 不存在")
	}

	delete(m1, "dema")
	s, ok = m1["dema"]
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("key 不存在")
	}
}
