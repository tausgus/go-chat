package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var users = make(map[net.Addr]string)
var activeConnections []net.Conn

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <port>")
		os.Exit(1)
	}

	listener, _ := net.Listen("tcp", ":"+os.Args[1])
	for {
		conn, _ := listener.Accept()
		fmt.Println("Accepting incoming connection from " + conn.RemoteAddr().String())
		go handleClient(conn) // Hand off any new incoming connections to handleClient
	}

}

func handleClient(client net.Conn) {
	var isConnected bool = true

	client.Write([]byte("Enter a username > "))
	message, _ := bufio.NewReader(client).ReadString('\n') // Promt the new client for an username and map it to their IP address
	users[client.RemoteAddr()] = sanitise(string(message))
	client.Write([]byte("Welcome, " + strings.TrimSuffix(users[client.RemoteAddr()], "\n") + ".\n"))

	activeConnections = append(activeConnections, client) // Append new connection to the list of active ones, to be used for broadcasting messages

	broadcast(client, "[SERVER] "+strings.TrimSuffix(users[client.RemoteAddr()], "\n")+" joined the room.\n")

	for isConnected {
		message, _ := bufio.NewReader(client).ReadString('\n')

		if message == "" { // Empty messages (no newlines) are sent by disconnected clients, this code will catch them and end the subroutine
			for _, connection := range activeConnections {
				isConnected = false
				client.Close()
				connection.Write([]byte("[SERVER] " + strings.TrimSuffix(users[client.RemoteAddr()], "\n") + " left the room.\n"))
			}
		} else {
			broadcast(client, "["+strings.TrimSuffix(users[client.RemoteAddr()], "\n")+"]: "+message) // If the message was not empty, broadcast it
		}
	}
}

func sanitise(s string) string {
	sanitised := ""
	blacklist := []string{"^[", "\033", "\u001b", "\x1b"}
	for _, value := range blacklist {
		sanitised = strings.ReplaceAll(s, value, "")
	}
	return sanitised
}

func broadcast(sender net.Conn, message string) {
	for _, connection := range activeConnections {
		if connection != sender { // Broadcast any new messages to all clients/connections except the sender
			connection.Write([]byte(sanitise(message)))
		}
	}
}
