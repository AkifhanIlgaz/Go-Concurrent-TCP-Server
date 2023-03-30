package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

const (
	MIN = 1
	MAX = 100
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConnectionHTTP(conn)
	}
}

func handleConnectionHTTP(connection net.Conn) {
	defer connection.Close()
	request(connection)
	response(connection)
}

func request(connection net.Conn) {
	i := 0
	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			m := strings.Fields(line)[0]
			fmt.Println("Methods", m)
		}

		if line == "" {
			break
		}
		i++
	}
}

func response(connection net.Conn) {
	body := `
This Is Go Http Server Using TCP
Golang HTTP Response
`

	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection, body)
}

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
