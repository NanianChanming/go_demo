package main

import "fmt"

/*func main() {
	u1 := User{"zhangsan", "student"}
	u2 := User{"lisi", "teacher"}
	u1.getName()
	u1.getWork()

	u2.getName()
	u2.getWork()

}*/

/*
go接口定义
Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}
*/
type getUserInfo interface {
	getName() string
	getWork() string
}

type User struct {
	name string
	work string
}

func (u User) getName() {
	fmt.Println(u.name)
}

func (u User) getWork() {
	fmt.Println(u.work)
}
