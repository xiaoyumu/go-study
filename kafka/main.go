package main

import (
	"log"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	broker     = kingpin.Flag("broker", "The kafka broker list.").Required().String()
	topic      = kingpin.Flag("topic", "The target topic to deliver message to.").Required().String()
	bindIP     = kingpin.Flag("bind-ip", "The IP address of the interface to bind to.").Default("0.0.0.0").String()
	listenPort = kingpin.Flag("listen-port", "The tcp port to listen for incoming clients.").Default("8092").Int()
	debug      = kingpin.Flag("debug", "Enable debug mode.").Bool()
)
 
 
func main() {

	kingpin.Version("1.0.0")
	kingpin.Parse()
	if *debug {
		log.Println("Debug Mode enabled.")
	}
	kafkaSetting := NewSettings(broker, topic, *debug)
	dispatcher := NewKafkaDispatcher(kafkaSetting)
	if err := dispatcher.Connect(); err != nil{
		panic(err)
	}
	log.Println("Dispatcher created.")
	proxy := NewTraceDataProxy(dispatcher)

	log.Println("TraceDataProxy created.")
	tcpServer := NewTCPServer(*bindIP, *listenPort, proxy)


	tcpServer.Startup()
}
