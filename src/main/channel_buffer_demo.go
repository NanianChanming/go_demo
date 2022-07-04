package main

import (
	"fmt"
)

/*func main() {
	ch := make(chan int)
	go write(ch)
	time.Sleep(2 * time.Second)

	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}*/

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("successfully wrote", i, "to ch")
		ch <- i
	}
	close(ch)
}
