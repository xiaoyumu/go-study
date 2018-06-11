package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/xiaoyumu/go-study/commandline"
)

func main() {
	pool, err := commandline.CreateDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	//pool.DumpParameters()

	clockHosts := make([]string, pool.Count())
	var intValue int
	index := &intValue

	fmt.Printf("Len: %d cap: %d\n", len(clockHosts), cap(clockHosts))

	pool.Iterate(func(name string, value string) {
		if !strings.HasPrefix(name, "clock") {
			return
		}

		// The slice and the value was enclosured in this func
		clockHosts[*index] = value

		// The address of index was enclosured in this func
		(*index)++
	})

	for _, value := range clockHosts {
		fmt.Println(value)
		indexOfDelimiter := strings.Index(value, "=")
		clockName := value[0 : indexOfDelimiter-1]
		clockAddress := value[indexOfDelimiter+1:]
		go showClock(clockName, clockAddress)
	}

	for {
		time.Sleep(100)
	}
}

func streamCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func showClock(name string, address string) {
	fmt.Printf("Connecting to clock %s at %s ...", name, address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	go streamCopy(os.Stdout, conn)
	streamCopy(conn, os.Stdin)
}
