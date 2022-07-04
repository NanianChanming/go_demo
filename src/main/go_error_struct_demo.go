package main

import "fmt"

func main() {
	//调用计算矩形面积的函数
	length, width := -9.0, -5.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*area2Error); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero, original err msg: %s", err.length, err.err)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero, original err msg: %s", err.width, err.err)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect ", area)
}

// 使用结构体类型的方法来提供更多错误信息
type area2Error struct {
	err    string
	length float64
	width  float64
}

//实现error接口，并给该错误类型添加两个方法
func (a *area2Error) Error() string {
	return a.err
}

//下面两个方法中，当length小于0时，lengthNegative返回true,而当width小于0时，widthNegative方法返回true
//这两个方法都提供了关于错误的更多信息，在这里它提示我们计算面积失败的原因（长度或宽度为负数），于是我们就有了两个错误类型结构体的方法，来提供更多的错误信息
func (a *area2Error) lengthNegative() bool {
	return a.length < 0
}

func (a *area2Error) widthNegative() bool {
	return a.width < 0
}

//计算面积的函数
//函数检查了长或宽是否小于0，如果小于0，rectArea会返回一个错误信息，否则会返回矩形面积和一个值为nil的错误
func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err += "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &area2Error{err, length, width}
	}
	return length * width, nil
}
