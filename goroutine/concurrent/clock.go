package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/xiaoyumu/go-study/commandline"
)

// DefaultHost for clock server to listen
const DefaultHost string = "localhost"

// DefaultPort for clock server to listen
const DefaultPort string = "8000"

func main() {
	pool := commandline.CreateDefault()
	clock(getClockSetting(pool))
}

func clock(name string, host string, port string) {
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

		go handleConnection(name, conn)
	}
}

func handleConnection(name string, conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	defer conn.Close()
	defer func() { log.Printf("Connection from %s was closed.", remoteAddr) }()
	for {
		log.Printf("Sending back current time to %s for clock [%s]", remoteAddr, name)
		_, err := io.WriteString(conn, fmt.Sprintf("[%s] %s\n", name, time.Now().Format(time.RFC3339)))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

const parameterNameHost string = "host"
const parameterNamePort string = "port"

func getClockSetting(pool *commandline.ParameterPool) (string, string, string) {
	return pool.GetParameterValueString("name", "Clock-"+os.Getenv("TZ")),
		pool.GetParameterValueString(parameterNameHost, DefaultHost),
		pool.GetParameterValueString(parameterNamePort, DefaultPort)
}
