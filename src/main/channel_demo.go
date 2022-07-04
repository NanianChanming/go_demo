package main

import "fmt"

/*func main() {
	// intSum()
	// strChannel()
	channel := bufferChannel(10)
	msgTest(channel)
	v, ok := <-channel
	if !ok {
		close(channel)
	} else {
		fmt.Println(v)
	}
	go fibonacci2(cap(channel), channel)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个数据之后就关闭了通道，
	// 所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了。
	for value := range channel {
		fmt.Println(value)
	}
}*/

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/*
Go 遍历通道与关闭通道
Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：
v, ok := <-ch
*/

/*
通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
ch := make(chan string 10)
带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
*/

func msgTest(channel chan string) {
	var s string
	for true {
		fmt.Scan(&s)
		if s == "exit" {
			break
		}
		sendChannel(s, channel)
	}
	// close(channel)
	// getChannel(channel)
}

func bufferChannel(len int) chan int {
	ch := make(chan int, len)
	return ch
}

func sendChannel(msg string, ch chan string) {
	ch <- msg
}

func getChannel(ch chan string) {
	s := <-ch
	fmt.Printf("from channel: %s\n", s)
}

func strChannel() {
	str := []string{"德玛", "西亚"}
	str1 := []string{"皇", "子"}
	// 定义一个通道
	ch := make(chan string)
	go appendStr(str1, ch)
	go appendStr(str, ch)
	s1, s2 := <-ch, <-ch
	fmt.Printf("%s %s", s1, s2)
}

func intSum() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 17
	go sum(s[len(s)/2:], c) // -5
	x, y := <-c, <-c        // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

/*
通道（channel）是用来传递数据的一个数据结构。
通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据 并把值赋给 v
*/

// 声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：
// ch := make(chan int)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func appendStr(arr []string, c chan string) {
	var str string
	for _, s := range arr {
		str += s
	}
	c <- str
}
