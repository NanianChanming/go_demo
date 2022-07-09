package main

import (
	"fmt"
	"os"
)

func main() {
	//writeStr()
	//writeBytes()
	//writeStrRow()
	appendStrRow()
}

// 将字符串写入文件
// 需要先创建文件，然后再将字符串写入文件
func writeStr() {
	f, err := os.Create("D:/GoProjects/hello_go/src/resource/test_write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("这是一首简单的小情歌")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 将字节写入文件
// 将字节写入文件和写入字符串非常类似，
// 使用Write方法将字节写入到文件
func writeBytes() {
	file, err := os.Create("D:/GoProjects/hello_go/src/resource/test_write_bytes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	b := []byte{100, 101, 102, 103, 104}
	n, err := file.Write(b)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(n, "byte[] written success")
}

// 将字符串一行一行的写入文件
func writeStrRow() {
	file, err := os.Create("D:/GoProjects/hello_go/src/resource/test_write_str_row.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d := []string{"这是一首简单的小情歌", "唱着我们心头的白鸽", "我想我很快乐", "当有你的温热"}
	for _, s := range d {
		_, err := fmt.Fprintln(file, s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written success")
}

// 追加到文件
// 将内容追加到刚才创建的test_write_str_row.txt 中
// 这个文件将以追加和写的方式打开，这些标志通过Open方法实现
func appendStrRow() {
	f, err := os.OpenFile("D:/GoProjects/hello_go/src/resource/test_write_str_row.txt", os.O_APPEND|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	newLine := "脚边的空气转了"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file append success")
}
