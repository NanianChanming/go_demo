package main

import (
	"fmt"
	"strings"
)

/*
类似策略模式
*/
func main() {
	s1 := student{"zhangfei", "B", "shu"}
	s2 := student{"guanyu", "A", "shu"}
	students := []student{s1, s2}
	// 筛选成绩为A的学生
	f := func(s student) bool {
		if s.grade == "A" {
			return true
		}
		return false
	}
	r := filter(students, f)
	fmt.Println(r)

	// 筛选姓zhang的学生
	f2 := func(s student) bool {
		if strings.Contains(s.name, "zhang") {
			return true
		}
		return false
	}
	r2 := filter(students, f2)
	fmt.Println(r2)
}

// 头等函数的实际用途
// 首先定义一个student类型
type student struct {
	name    string
	grade   string
	country string
}

// 编写一个filter函数，该函数接收一个students切片和一个函数作为参数
// filter方法第二个参数是一个函数，这个函数接收student参数，返回bool值
// 这个函数计算了某一学生是否满足筛选条件
func filter(students []student, f func(s student) bool) []student {
	var stus []student
	for _, v := range students {
		if f(v) {
			stus = append(stus, v)
		}
	}
	return stus
}
