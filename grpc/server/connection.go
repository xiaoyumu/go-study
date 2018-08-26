package main

import (
	"database/sql"
	"sync"
)

type SQLConnection struct {
	server                  ServerInfo
	currentConnectionString string
	referenceCount          int64
	maximumCount            int64
	mutex                   sync.RWMutex
}

type ServerInfo struct {
	Server   string
	Port     int32
	Database string
	UserID   string
	Password string
}

type SQLDB struct {
	ID int32
	DB *sql.DB
}

type Connection interface {
	// Initialize Connection
	Open() error
	Close() error

	// Aquire this connection, if the reference count within limits
	Aquire() (*SQLDB, error)

	// Release a connection, reduce the reference count.
	Release(*SQLDB) error
	// Returns the server info
	Server() ServerInfo

	// Returns the current reference count
	ReferenceCount() int64

	// Returns the maximum count of the connection allowed
	MaximumCount() int64
}

// Server function returns the ServerInfo of this connection
func (c *SQLConnection) Server() ServerInfo {
	return c.server
}

// ReferenceCount returns the current reference count
func (c *SQLConnection) ReferenceCount() int64 {
	return c.referenceCount
}

// Aquire new connection, if the active connection is full (=MaximumCount)
// No more SQLDB can be aquired, an error will be returned to the caller
func (c *SQLConnection) Aquire() (*SQLDB, error) {
	return nil, nil
}

func (c *SQLConnection) Release(db *SQLDB) error {
	return nil
}

func (c *SQLConnection) Open() error {
	return nil
}

func (c *SQLConnection) Close() error {

	return nil
}
