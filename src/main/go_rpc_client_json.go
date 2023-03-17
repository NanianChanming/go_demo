package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith3.Multiply", args, &reply)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	fmt.Printf("Arith3: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith3.Divide", args, &quot)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	fmt.Printf("Arith3: %d / %d = %d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
