package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	listenr, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}
	defer listenr.Close()

	conn, err := listenr.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established from", conn.RemoteAddr())

	ch := getLinesChannel(conn)

	for line := range ch {
		fmt.Println(line)
	}

	fmt.Println("Connection closed from", conn.RemoteAddr())
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		defer f.Close()

		var curLine string
		buf := make([]byte, 8)
		for {
			n, err := f.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				panic(err)
			}
			if n == 0 {
				break
			}
			line := string(buf[0:n])
			curLine += line

			result := strings.Split(curLine, "\n")

			if len(result) > 1 {
				for i := 0; i < len(result)-1; i++ {
					ch <- result[i]
				}
				curLine = result[len(result)-1]
			}
		}
		if len(curLine) > 0 {
			ch <- curLine
		}
	}()
	return ch
}
