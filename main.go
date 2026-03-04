package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		fmt.Printf("Error getting free port: %v\n", err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("Error getting free port: %v\n", err)
	}
	defer l.Close()

	port := l.Addr().(*net.TCPAddr).Port

	fmt.Printf("%d\n", port)
}
