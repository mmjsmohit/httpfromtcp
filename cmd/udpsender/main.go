package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, addrErr := net.ResolveUDPAddr("udp", "localhost:42069")
	if addrErr != nil {
		log.Fatal("Error connecting to address")
		return
	}

	conn, connErr := net.DialUDP("udp", nil, addr)

	if connErr != nil {
		log.Fatal("Error establishing connection")
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		userInput, inputErr := reader.ReadString('\n')

		if inputErr != nil {
			log.Fatal("Error while taking user input")
			return
		}

		conn.Write([]byte(userInput))
	}
}
