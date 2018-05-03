package dboperator

import "fmt"

// MSSQLConnection represent a MS SQL Server connection
type MSSQLConnection struct {
	ci    *ConnectionInfo // 从ConnectionInfo 获得 SetConnectionString的实现
	state int8
}

// Open a MSSQLConnection
func (c *MSSQLConnection) Open() {
	fmt.Printf("Connected to MS SQL Server (%s)\r\n", c.ci.GetConnectionString())
}

// Close a MSSQLConnection
func (c *MSSQLConnection) Close() {
	fmt.Println("Connection to MS SQL Server has been closed.")
}

// GetState returns current state of a MSSQLConnection
func (c *MSSQLConnection) GetState() int8 {
	return c.state
}
