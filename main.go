package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	ch := getLinesChannel(file)

	for {
		line, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("read: %s\n", line)
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
