package main

import "fmt"

/*func main() {
	// useFor()
	MultiTable()
}*/

func MultiTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}

func useFor() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 100 {
		fmt.Println(i)
		i++
	}
}
