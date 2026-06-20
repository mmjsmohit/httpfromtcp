package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	file, err := os.Open("messages.txt")

	if err != nil {
		fmt.Println("An error occured while reading the file")
		return
	}

	returnedChan := getLinesChannel(file)

	for i := range returnedChan {
		fmt.Println("read: " + i)
	}

}
