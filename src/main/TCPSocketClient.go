package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:post ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	fmt.Println("service = ", service)
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	//checkError(err)
	fmt.Println("tcpAddr = ", tcpAddr)
	for {
		conn, _ := net.DialTCP("tcp", nil, tcpAddr)
		conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
		result, _ := ioutil.ReadAll(conn)
		fmt.Println(string(result))
		conn.Close()
	}

}
