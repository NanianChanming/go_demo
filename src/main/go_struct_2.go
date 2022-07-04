package main

import "fmt"

func main() {
	a1 := author{"云长", "关", "three kingdom"}
	a1.fullName()
	p1 := post{"abc1", "abcdefg1", a1}
	p2 := post{"abc2", "abcdefg2", a1}
	p3 := post{"abc3", "abcdefg3", a1}
	p4 := post{"abc4", "abcdefg4", a1}
	w := website{post: []post{p1, p2, p3, p4}}
	w.contents()
}

// author 组合代替继承
type author struct {
	firstname string
	lastname  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstname, a.lastname)
}

// 创建post结构体
type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("title: ", p.title)
	fmt.Println("content: ", p.content)
	// 一旦结构体内嵌套了一个结构体字段，Go 可以使我们访问其嵌套的字段，好像这些字段属于外部结构体一样。
	// 所以下面第 35 行的 p.author.fullName() 可以替换为 p.fullName()。
	// fmt.Println("Author: ", p.author.fullName())
	// 于是details() 方法可以重写，如下所示：
	fmt.Println("Author: ", p.fullName())
	fmt.Println("bio: ", p.author.bio)
}

// 结构体切片嵌套
// 创建一个切片结构体
type website struct {
	post []post
}

func (w website) contents() {
	fmt.Println("Contents of Website \n")
	for _, v := range w.post {
		v.details()
		fmt.Println()
	}
}
