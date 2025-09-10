package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:42069")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

		buf, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if _, err := conn.Write([]byte(buf)); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

}
