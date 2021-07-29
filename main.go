package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"os"
)


var users = make(map[net.Addr]string)
var activeConnections []net.Conn


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <port>")
		os.Exit(1)
	}

	listener, _ := net.Listen("tcp", ":" + os.Args[1])
	for {
		conn, _ := listener.Accept()
		fmt.Println("Accepting incoming connection from " + conn.RemoteAddr().String())
		go handleClient(conn)
	}
}


func handleClient(client net.Conn) {
	activeConnections = append(activeConnections, client) 

	client.Write([]byte("Enter a username > "))
        message, _ := bufio.NewReader(client).ReadString('\n')
        users[client.RemoteAddr()] = string(message)

	for {
		message, _ := bufio.NewReader(client).ReadString('\n')

		for i := 0; i < len(activeConnections); i++ {
			if activeConnections[i] != client {
				activeConnections[i].Write([]byte("[" + strings.TrimSuffix(users[client.RemoteAddr()], "\n") + "]: " + message))
			}
		}
	}
}
