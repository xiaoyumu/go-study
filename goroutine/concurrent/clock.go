package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	clock()
}

func clock() {
	log.Print("Starting single connection clock ...")
	host := "localhost"
	port := "8000"
	endpoint := host + ":" + port
	linstener, err := net.Listen("tcp", endpoint)
	log.Printf("Listening on %s", endpoint)
	log.Printf("Use nc %s %s to test connection.", host, port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := linstener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("Connection from %s was accepted.", conn.RemoteAddr())

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	defer func() { log.Print("Connection closed.") }()
	for {
		log.Print("Sending back current time.")
		_, err := io.WriteString(conn, time.Now().Format(time.RFC3339+"\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
