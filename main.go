package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:42069")
	if err != nil {
		listener.Close()
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		fmt.Println("Connection has been accepted")

		ch := getLinesChannel(conn)
		for {
			line, ok := <-ch
			if !ok {
				break
			}
			fmt.Printf("%s\n", line)
		}
		fmt.Println("Connection has been closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		line := ""
		for {
			buf := make([]byte, 8)
			if _, err := f.Read(buf); err == io.EOF {
				ch <- line
				break
			}

			bufString := strings.Split(string(buf), "\n")
			line += bufString[0]

			if len(bufString) > 1 {
				ch <- line
				line = bufString[1]
			}
		}
		close(ch)
		f.Close()
	}()

	return ch
}
