package dboperator

import "fmt"

// MySQLConnection represent a  MySQL Server connection
type MySQLConnection struct {
	ci    *ConnectionInfo // 从ConnectionInfo 获得 SetConnectionString的实现
	state int8
}

// Open a MySQLConnection
func (c *MySQLConnection) Open() {
	fmt.Printf("Connected to MySQL Server (%s)\r\n", c.ci.GetConnectionString())
}

// Close a MySQLConnection
func (c *MySQLConnection) Close() {
	fmt.Println("Connection to MySQL Server has been closed.")
}

// GetState returns current state of a MySQLConnection
func (c *MySQLConnection) GetState() int8 {
	return c.state
}
