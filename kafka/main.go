package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/tidwall/evio"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// The variables are pointer of the value, to access the value instread of the pointer,
	// * need to be added, for example: *bindIP gets the actual string value of the variable.
	bindIP     = kingpin.Flag("bind-ip", "The ip address to bind to.").Default("0.0.0.0").Short('b').String()
	port       = kingpin.Flag("port", "The port to listen to.").Default("8092").Short('p').Int()
	responseOK = []byte("OK\n")
)

func main() {

	kingpin.Version("1.0.0")
	kingpin.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	var events evio.Events
	events.Data = func(c evio.Conn, in []byte) ([]byte, evio.Action) {
		resp, err := ProcessTraceMessage(in)
		if err != nil {
			return []byte(fmt.Sprintln(err)), evio.Close
		}
		return resp, evio.None
	}

	events.Opened = func(c evio.Conn) (out []byte, options evio.Options, action evio.Action) {
		log.Printf("Client %s connected", c.RemoteAddr())
		return
	}

	events.Closed = func(c evio.Conn, err error) (action evio.Action) {
		if err != nil {
			log.Printf("Client %s closed with err %s", c.RemoteAddr(), err)
			return
		}

		log.Printf("Client %s closed ", c.RemoteAddr())
		return
	}

	events.Serving = func(srv evio.Server) (action evio.Action) {
		log.Printf("TCP server started on port %d (loops: %d)", port, srv.NumLoops)
		return
	}

	log.Printf("Start to bind %s:%d", *bindIP, *port)

	if err := evio.Serve(events, fmt.Sprintf("tcp://%s:%d", *bindIP, *port)); err != nil {
		panic(err.Error())
	}

	log.Println("Shuting down.")
}

// ProcessTraceMessage handles business logic.
func ProcessTraceMessage(inData []byte) ([]byte, error) {
	log.Printf("Message received.")
	req, err := DeserializeTraceMessage(inData)

	if err != nil {
		return nil, fmt.Errorf("invalid request, %s", err)
	}

	if req.Key == nil {
		return nil, errors.New("invalid request, The trace message property Key is nil")
	}

	if req.Value == nil {
		return nil, errors.New("invalid request, The trace message property Value is nil")
	}

	log.Printf("Key: %s", string(*req.Key))
	log.Printf("Value: %s", string(*req.Value))

	return responseOK, nil
}
