package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func main() {
	//writeStr()
	//writeBytes()
	//writeStrRow()
	//appendStrRow()
	writeWithChannel()
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

// 并发写文件
// 当多个goroutines同时并发写文件时，我们会遇到竞争条件，因此，当发生同步写的时候，需要一个channel作为一致写入的条件
/**
写一个方法创建100个goroutines，每个goroutine将并发产生一个随机数，这些随机数将被写入到文件里面
1.创建一个channel用来读和写这个随机数
2.创建100个生产者goroutine，每个goroutine都产生一个随机数并写入到channel里
3.创建一个消费者goroutine用来从channel读取随机数并写入文件，这样的话，就只有一个goroutine向文件中写数据，从而避免竞争条件
4.完成之后关闭文件
*/
// 产生随机数的方法
func produce(data chan int, wg *sync.WaitGroup) {
	// 产生一个999以内的随机数
	n := rand.Intn(999)
	// 写入到管道
	data <- n
	wg.Done()
}

// 写文件的函数
// 这个consume方法创建了一个名为test_write_channel.txt的文件，然后从channel中读取随机数并写到文件中，
// 一旦读取完成并且将随机数写入文件后，通过done这个channel中写入true来通知任务已经完成
func consume(data chan int, done chan bool) {
	f, err := os.Create("D:/GoProjects/hello_go/src/resource/test_write_channel.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err := fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

// 提供完成程序，并发写入文件
func writeWithChannel() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d {
		fmt.Println("file written success")
	} else {
		fmt.Println("file written failed")
	}
}
