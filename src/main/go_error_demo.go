package main

import "fmt"

/*func main() {
	i, msg := Divide(100, 10)
	fmt.Println(i, msg)
	i, msg = Divide(100, 0)
	if msg != "" {
		fmt.Println(msg)
	}

}*/

/*
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
error类型是一个接口类型，这是它的定义：
type error interface {
	Error() string
}
*/

/*
我们可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
*/

// DivideError 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现Error接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0`
	return fmt.Sprintf(strFormat, de.dividee)
}

// Divide 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	}
	return varDividee / varDivider, ""
}
