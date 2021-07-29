package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"os"
)


var users = make(map[net.Addr]string)
var conns []net.Conn

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <port>")
		os.Exit(1)
	}
	listener, _ := net.Listen("tcp", ":" + os.Args[1])
	for {
		conn, _ := listener.Accept()
		fmt.Println("Proceeding to handle a connection..")

		go handleConn(conn)
	}
}

func handleConn(workConn net.Conn) {
	defer workConn.Close()
	conns = append(conns, workConn)
	workConn.Write([]byte("Enter your name: "))
        message, _ := bufio.NewReader(workConn).ReadString('\n')
        users[workConn.RemoteAddr()] = string(message)
	workConn.Write([]byte("Welcome, " + users[workConn.RemoteAddr()]))

	for {
		message, _ := bufio.NewReader(workConn).ReadString('\n')
		for i := 0; i < len(conns); i++ {
			if conns[i] != workConn {
				conns[i].Write([]byte("[" + strings.TrimSuffix(users[workConn.RemoteAddr()], "\n") + "]: " + message))
			}
		}
	}
}
