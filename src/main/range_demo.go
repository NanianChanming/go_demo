package main

import "fmt"

/*func main() {
	ArrRange()
}*/

/*
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
*/

func ArrRange() {
	arr := []string{"德玛", "西亚", "皇子"}
	for _, s := range arr {
		fmt.Printf("%s \n", s)
	}
	fmt.Println("---------")

	as := arr[:2]
	for _, s := range as {
		fmt.Printf("%s \n", s)
	}
	fmt.Println("---------")

	map1 := make(map[string]string)
	map1["诺克"] = "萨斯"
	map1["扭曲"] = "丛林"
	for k, v := range map1 {
		fmt.Printf("key = %s, value = %s \n", k, v)
	}

}
