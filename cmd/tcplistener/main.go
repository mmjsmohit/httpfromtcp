package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {

	stringChan := make(chan string)

	dataRemaining := true

	currentLine := ""

	go func() {
		defer close(stringChan)

		for dataRemaining {
			dataBytes := make([]byte, 8)
			n, _ := f.Read(dataBytes)
			parts := bytes.Split(dataBytes, []byte("\n"))
			if len(parts) == 2 {
				currentLine += string(parts[0])
				// fmt.Println("read: " + currentLine) Send the string to the channel
				stringChan <- currentLine
				currentLine = string(parts[1])
			} else {
				currentLine = currentLine + string(dataBytes)
			}

			if n == 0 {
				dataRemaining = false
				f.Close()
				return
			}
		}
	}()

	return stringChan
}

func main() {
	listener, listenErr := net.Listen("tcp", ":42069")

	if listenErr != nil {
		fmt.Errorf("Error encountered while creating a listener")
		return
	}

	// close the listener after the program exits
	defer fmt.Println("Connection closed")
	defer listener.Close()

	for {
		conn, connErr := listener.Accept()
		if connErr != nil {
			fmt.Errorf("Error while accepting connection")
			return
		}
		fmt.Println("Connection accepted")

		returnedChan := getLinesChannel(conn)
		for i := range returnedChan {
			fmt.Println("read: " + i)
		}
	}
}
