package main

import "fmt"

// Income
// 使用接口实现多态
// 首先定义一个接口
type Income interface {
	calculate() int
	source() string
}

// FixedBilling 定义结构体
type FixedBilling struct {
	// 项目名称
	projectName string
	// 投标金额
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

// 分别给两个结构体定义方法
// 由于 FixedBilling 和 TimeAndMaterial 两个结构体都定义了 Income 接口的两个方法：calculate() 和 source()，因此这两个结构体都实现了 Income 接口。
func (f FixedBilling) calculate() int {
	// 在项目 FixedBilling 里面，收入就是项目的投标金额。因此我们返回 FixedBilling 类型的calculate() 方法返回值。
	return f.biddedAmount
}
func (f FixedBilling) source() string {
	return f.projectName
}
func (t TimeAndMaterial) calculate() int {
	// 而在项目 TimeAndMaterial 里面，收入等于 noOfHours 和 hourlyRate 的乘积，作为 TimeAndMaterial 类型的 calculate() 方法返回值。
	return t.noOfHours * t.hourlyRate
}
func (t TimeAndMaterial) source() string {
	return t.projectName
}

// 定义一个calculateNetIncome函数用来计算并打印总收入。
func calculateNetIncome(ins []Income) int {
	netIncome := 0
	for _, v := range ins {
		fmt.Printf("Income From %s = $%d\n", v.source(), v.calculate())
		netIncome += v.calculate()
	}
	fmt.Printf("Net income of organisation = $%d \n", netIncome)
	return netIncome
}

func main() {
	f := FixedBilling{"三国", 100}
	t := TimeAndMaterial{"水浒", 2, 50}
	ins := []Income{f, t}
	fmt.Println(calculateNetIncome(ins))
}
