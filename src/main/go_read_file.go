package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//readTest()
	//fptr := flag.String("fpath", "test.txt", "file path to read from")
	//flag.Parse()
	//fmt.Println("value of fpath is", *fptr)
	//readTestFlag(fptr)
	//readTestBuffer()
	readTestScan()
}

/*
读取文件
将文件读取到内存是最基本的操作之一，这需要使用到ioutil包中的ReadFile函数
在src/resource目录下创建test.txt，然后来读取它
如果程序位置与文件所在位置不同，会打印出下面的错误：
File reading error:  open ../resource/test.txt: The system cannot find the path
这是因为go是编译型语言，go文件执行时会创建一个二进制文件。二进制文件独立于源代码，可以在任何位置运行。
由于在运行二进制文件的位置上没有找到test.txt,因此程序会报错，提示无法找到指定的文件
有三种方法可以解决这个问题
1.使用绝对文件路径
2.使用命令行标记来传递文件路径（未实测）
3.将文件绑定在二进制文件中（暂不研究）
*/
func readTest() {
	// 绝对路径
	file, err := ioutil.ReadFile("D:/GoProjects/hello_go/src/resource/test.txt")
	if err != nil {
		fmt.Println("File reading error: ", err)
		return
	}
	fmt.Println("file contents: ", string(file))
}

// 使用命令行标记来传递文件路径(未实测)
func readTestFlag(fptr *string) {
	// 绝对路径
	file, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error: ", err)
		return
	}
	fmt.Println("file contents: ", string(file))
}

// 分块读取文件
// 当文件非常大时，尤其在RAM存储量不足的情况下，把整个文件都读取到内存中是没有意义的。更好的方法是分块读取文件
// 这可以使用bufio包来完成
// 下面这个方法，以3个字节的块为单位读取test.txt文件
func readTestBuffer() {
	f, err := os.Open("D:/GoProjects/hello_go/src/resource/test.txt")
	if err != nil {
		fmt.Println("File reading error: ", err)
		return
	}
	reader := bufio.NewReader(f)
	b := make([]byte, 4)
	for {
		_, err := reader.Read(b)
		if err != nil {
			fmt.Println("reading file error: ", err)
			break
		}
		fmt.Println(string(b))
	}
}

//逐行读取
func readTestScan() {
	f, err := os.Open("D:/GoProjects/hello_go/src/resource/test.txt")
	if err != nil {
		fmt.Println("File reading error: ", err)
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	error := scanner.Err()
	if err != nil {
		log.Fatal(error)
	}
}
