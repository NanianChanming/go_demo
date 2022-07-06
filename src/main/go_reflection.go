package main

import (
	"fmt"
	"reflect"
)

/*
代码中应该使用反射吗
引用Rob Pike关于反射使用的格言： 清晰优于聪明，而反射并不是一目了然的
反射是go语言中非常强大和高级的概念，应该小心谨慎的使用它。
使用反射编写清晰和可维护的代码是很困难的，应该尽量避免使用。
*/
func main() {
	//s := createSqlQuery(getOrder())
	//fmt.Println(s)
	//reflectDemo(getOrder())
	//reflectDemo2(getOrder())
	//reflectDemo3()
	//createQuery(getOrder())
	//createQuery(getUser())
	createQuery1(getOrder())
	createQuery1(getUser())
}

// 反射
// 反射是指程序能够在运行时检查变量和值，求出他们的类型

//此方法中，i的类型在编译时就知道了
func getIntType() {
	i := 1
	fmt.Printf("%d, %T", i, i)
}

//下面的方法中，接收一个结构体作为参数，并用它来创建一个sql插入查询
type order struct {
	orderId    int
	customerId int
}

func getOrder() order {
	var o order
	o = order{
		orderId:    1234,
		customerId: 5678,
	}
	return o
}

// 下面函数会返回 insert into order values(1234, 5678)
func createSqlQuery(o order) string {
	return fmt.Sprintf("insert into order values(%d, %d)", o.orderId, o.customerId)
}

/*
现在来升级这个方法，让它变的更通用。
由于设想该函数可以适用于任何结构体类型，所以接收interface{}作为参数。
*/
type user struct {
	id     int
	name   string
	age    int
	gender string
}

func getUser() user {
	return user{
		id:     1,
		name:   "zhangfei",
		age:    18,
		gender: "male",
	}
}

/**
在go中，reflect实现了运行时反射，reflect包会帮助识别interface{}变量底层具体类型和具体值。
reflect.Type和reflect.Value
代码输出:
type  main.order
value  {1234 5678}

reflect包下另一个类型
reflect.kind 代码输出: kind  struct
Type表示interface{}的实际类型（main.Order）
kind则表示该类型的特定类别（struct）

NumField()和Field()方法
NumField()方法返回结构体中字段的数量，
而Field(i int)方法返回字段i的reflect.Value
*/
func reflectDemo(i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := t.Kind()
	fmt.Println("type ", t)
	fmt.Println("value ", v)
	fmt.Println("kind ", k)
}

/*
NumField()和Field()方法
NumField()方法返回结构体中字段的数量，
而Field(i int)方法返回字段i的reflect.Value
*/
func reflectDemo2(i interface{}) {
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		v := reflect.ValueOf(i)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

/*
Int()和String()方法
这两个方法可以帮助我们分别取出reflect.Value作为int64和string
*/
func reflectDemo3() {
	a := 2
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v \n", x, x)

	b := "guan yun chang"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v \n", y, y)
}

/*
完整实现
*/
func createQuery(i interface{}) string {
	if reflect.ValueOf(i).Kind() != reflect.Struct {
		return fmt.Sprintf("unsupported type")
	}
	t := reflect.TypeOf(i).Name()
	sql := fmt.Sprintf("insert into %s values (", t)
	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				sql = fmt.Sprintf("%s%d", sql, v.Field(i).Int())
			} else {
				sql = fmt.Sprintf("%s, %d", sql, v.Field(i).Int())
			}
		case reflect.String:
			if i == 0 {
				sql = fmt.Sprintf("%s\"%s\"", sql, v.Field(i).String())
			} else {
				sql = fmt.Sprintf("%s, \"%s\"", sql, v.Field(i).String())
			}
		default:
			fmt.Println("unsupported type")
			return "unsupported type"
		}
	}
	sql = fmt.Sprintf("%s)", sql)
	fmt.Println(sql)
	return sql
}

//向输出的查询中添加字段名，请尝试着修改程序，打印出以下格式的查询：
//insert into order(ordId, customerId) values(456, 56)
func createQuery1(obj interface{}) string {
	if reflect.ValueOf(obj).Kind() != reflect.Struct {
		return fmt.Sprintf("unsupported type")
	}
	t := reflect.TypeOf(obj).Name()
	sqlPre := fmt.Sprintf("insert into %s (", t)
	sqlSuf := fmt.Sprintf("values (")
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Int:
			if i == 0 {
				sqlPre = fmt.Sprintf("%s%s", sqlPre, reflect.TypeOf(obj).Field(i).Name)
				sqlSuf = fmt.Sprintf("%s%d", sqlSuf, v.Field(i).Int())
			} else {
				sqlPre = fmt.Sprintf("%s,%s", sqlPre, reflect.TypeOf(obj).Field(i).Name)
				sqlSuf = fmt.Sprintf("%s,%d", sqlSuf, v.Field(i).Int())
			}
		case reflect.String:
			if i == 0 {
				sqlPre = fmt.Sprintf("%s%s", sqlPre, reflect.TypeOf(obj).Field(i).Name)
				sqlSuf = fmt.Sprintf("%s\"%s\"", sqlSuf, v.Field(i).String())
			} else {
				sqlPre = fmt.Sprintf("%s,%s", sqlPre, reflect.TypeOf(obj).Field(i).Name)
				sqlSuf = fmt.Sprintf("%s,\"%s\"", sqlSuf, v.Field(i).String())
			}
		default:
			fmt.Println("unsupported type")
			return "unsupported type"
		}
	}
	var sql = fmt.Sprintf("%s) %s)", sqlPre, sqlSuf)
	fmt.Println(sql)
	return sql
}
