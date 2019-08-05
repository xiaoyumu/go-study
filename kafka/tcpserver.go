package main

import (
	"fmt"
	"github.com/tidwall/evio"
	"log"
)

// TCPServer interface defines the basic behaviors of a TCPServer
type TCPServer interface{
	Startup()
}

const (
	StatusAccepted    = "202"
	StatusBadRequest  = "400"
	StatusServerError = "500"
)

type tcpServer struct {
	listenIP string
	port int
	events *evio.Events
	handler TraceDataHandler
	shutdown bool
}

// NewTCPServer function creates an instance of TCPServer
func NewTCPServer(listenIP string, port int, handler TraceDataHandler) TCPServer {
	server := &tcpServer{
		listenIP: listenIP,
		port: port,
		handler: handler,
		events: new(evio.Events),
	}

	server.events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		return server.OnDataReceived(c, in)
	}

	server.events.Opened = func(c evio.Conn) (out []byte, options evio.Options, action evio.Action) {
		log.Printf("Client %s connected", c.RemoteAddr())
		return
	}

	server.events.Closed = func(c evio.Conn, err error) (action evio.Action) {
		if err != nil {
			log.Printf("Client %s closed with err %s", c.RemoteAddr(), err)
			return
		}

		log.Printf("Client %s closed ", c.RemoteAddr())
		return
	}

	server.events.Serving = func(srv evio.Server) (action evio.Action) {
		log.Printf("TCP server started on port %d (loops: %d)", port, srv.NumLoops)
		return
	}

	return server
}

func (svr *tcpServer) OnDataReceived(c evio.Conn, in []byte) ([]byte, evio.Action){
	resp := TraceResponse{
		Status: StatusAccepted,
	}
	req, err := Decode(in)
	if err != nil {
		resp.Status = StatusBadRequest
		resp.Message = err.Error()
		return resp.Serialize(), evio.Close
	}

	err = svr.handler.ProcessTraceData(req)
	if err != nil {
		resp.Status = StatusServerError
		resp.Message = err.Error()
		return resp.Serialize(), evio.Close
	}

	return resp.Serialize(), evio.None
}

func (svr *tcpServer) Startup() {
	log.Printf("Start to bind %s:%d", svr.listenIP, svr.port)
	log.Println("*********** Press Ctrl + C to quit ***********")
	if err := evio.Serve(*svr.events, fmt.Sprintf("tcp://%s:%d", svr.listenIP, svr.port)); err != nil {
		panic(err.Error())
	}
}
