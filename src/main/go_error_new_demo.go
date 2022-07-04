package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	radius := -20.0
	//area, err := circleArea(radius)
	/*if err != nil {
		fmt.Println(err)
		return
	}*/
	area, err := circleAreaErrorStruct(radius)
	if err != nil {
		if e, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero: %s", e.radius, e.err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("area of circle %0.2f", area)
}

// 使用new函数创建自定义错误
// 创建自定义错误最简单的方法是使用errors包中的new函数
// 创建一个计算圆面积的方法，如果为负，则会返回一个错误
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

//使用Errorf函数给错误添加更多的信息
//需要用到fmt包，Errorf函数会根据格式说明符，规定错误的格式，并返回一个符合该错误的字符串
func circleAreaErrorf(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

//使用结构体和字段提供错误的更多信息
//错误还可以用实现了error接口的结构体来表示，这种方式可以更加灵活的处理错误
//定义一个表示错误的结构体类型，错误类型的命名约定是名称以Error结尾
type areaError struct {
	err    string
	radius float64
}

//实现error接口
func (a *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", a.radius, a.err)
}

// 改写方法
func circleAreaErrorStruct(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative ", radius}
	}
	return math.Pi * radius * radius, nil
}
