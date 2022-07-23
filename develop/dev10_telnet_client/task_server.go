package main

import (
	"log"
	"net"
	"time"
)
import "fmt"
import "bufio"
import "strings"

func myHandle(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		switch message {
		case "time is up":
			return
		default:
			fmt.Print("Message Received:", message)
			newMessage := strings.ToUpper(message)
			_, err := conn.Write([]byte(time.Now().String() + " " + newMessage + "\n"))
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:3000")
	conn, _ := ln.Accept()

	myHandle(conn)

	err := ln.Close()
	if err != nil {
		log.Print(err)
	}
}
