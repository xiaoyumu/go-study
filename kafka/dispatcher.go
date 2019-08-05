package main


// Message contains the message content to be dispatched
type Message struct {
	Key   []byte
	Value []byte
}

// Dispatcher interface defines the basic operation for
// a message dispatcher
type Dispatcher interface {
	// Connect to the dispatch endpoint
	Connect() error

	// Dispatch a message entity
	DispatchMessage(msg *Message) error

	// Dispatch data by given key and value
	Dispatch(key []byte, value []byte) error

	// Close the connection to kafka brokers
	Close()
}
