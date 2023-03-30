package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	port := ":" + arguments[1]

	initiateConcurrentTCP(port)
}

const (
	MIN = 1
	MAX = 100
)

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(connection net.Conn) {
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		netData = strings.TrimSpace(netData)
		if netData == "STOP" {
			break
		}

		response := strconv.Itoa(random()) + "\n"
		connection.Write([]byte(response))
	}

	connection.Close()
}

func initiateConcurrentTCP(port string) {
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(connection)
	}
}
