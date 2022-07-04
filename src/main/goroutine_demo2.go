package main

import (
	"fmt"
	"time"
)

func printStr() {
	var times int
	for {
		times++
		fmt.Println(times)
		// 延时1秒
		time.Sleep(time.Second * 5)
	}
}

/*func main() {
	go printStr()

	var s string
	for {
		fmt.Scan(&s)
		fmt.Print(s)
	}
}*/
