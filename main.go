package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*

program should start by using net.ResolveUDPAddr to resolve the address localhost:42069
Use net.DialUDP to prepare a UDP connection, and defer the closing of the connection.
Create a new bufio.Reader that reads from os.Stdin
Start an infinite loop that:
Prints a > character to the console (to indicate that the program is ready for user input)
Reads a line from the bufio.Reader using reader.ReadString, and log any errors
Writes the line to the UDP connection using conn.Write, and log any errors
*/

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
