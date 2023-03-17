package main

import (
	"fmt"
	"net/rpc"
)

type Args1 struct {
	A, B int
}

type Quotient1 struct {
	Quo, Rem int
}

/*
基于TCP实现的RPC客户端
*/
func main() {
	client, err := rpc.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	args := Args1{17, 8}
	var reply int
	err = client.Call("Arith2.Multiply", args, &reply)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	fmt.Printf("Arith2: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient1
	err = client.Call("Arith2.Divide", args, &quot)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	fmt.Printf("Arith2: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
