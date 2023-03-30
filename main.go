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

const (
	MIN = 1
	MAX = 100
)

type Server struct {
	host string
	port string
}

type Client struct {
	conn net.Conn
}

type Config struct {
	Host string
	Port string
}

func New(config *Config) Server {
	return Server{
		host: config.Host,
		port: config.Port,
	}
}

func (srv Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", srv.host, srv.port))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		client := Client{conn}

		go client.HandleRequest()
	}
}

func (client Client) HandleRequest() {
	fmt.Printf("Serving %s\n", client.conn.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		netData = strings.TrimSpace(netData)
		if netData == "STOP" {
			break
		}

		response := strconv.Itoa(random()) + "\n"
		client.conn.Write([]byte(response))
	}

	client.conn.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number")
		return
	}

	cfg := Config{
		Port: arguments[1],
	}

	server := New(&cfg)
	server.Run()
}

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}
