package main

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/xiaoyumu/go-study/commandline"
)

// DefaultHost for clock server to listen
const DefaultHost string = "localhost"

// DefaultPort for clock server to listen
const DefaultPort string = "8000"

func main() {
	pool := commandline.Create()
	clock(getHostAndPort(pool))
}

func clock(host string, port string) {
	log.Print("Starting single connection clock ...")
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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	defer conn.Close()
	defer func() { log.Printf("Connection from %s was closed.", remoteAddr) }()
	for {
		log.Printf("Sending back current time to %s", remoteAddr)
		_, err := io.WriteString(conn, time.Now().Format(time.RFC3339+"\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

const parameterNameHost string = "host"
const parameterNamePort string = "port"

func getHostAndPort(pool *commandline.ParameterPool) (string, string) {
	return pool.GetParameterValueString(parameterNameHost, DefaultHost),
		pool.GetParameterValueString(parameterNamePort, DefaultPort)
}
