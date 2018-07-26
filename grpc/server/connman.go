package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

// ConnectionManager contains function definitions related to a connection pooling management
type ConnectionManager interface {
	// GetConnection returns a ptr to sql.DB based on given connection string
	GetConnection(connStr string) (*sql.DB, error)
}

// BasicConnectionManager is the basic connection pool implementation
type BasicConnectionManager struct {
	connectionPool map[string]*sql.DB
	mutex          sync.RWMutex
}

var (
	connectionManager ConnectionManager
)

func init() {
	connectionManager = &BasicConnectionManager{
		connectionPool: make(map[string]*sql.DB),
		mutex:          sync.RWMutex{},
	}
}

// GetConnectionManager function return the singleton instance of ConnectionManager
func GetConnectionManager() ConnectionManager {
	return connectionManager
}

// GetConnection returns a ptr to sql.DB based on given connection string
func (bcm *BasicConnectionManager) GetConnection(connStr string) (*sql.DB, error) {
	if len(connStr) == 0 {
		return nil, fmt.Errorf("the connStr parameter cannot be empty")
	}

	conn, ok := bcm.connectionPool[connStr]
	if ok {
		return conn, nil
	}

	bcm.mutex.Lock()
	defer bcm.mutex.Unlock()

	conn, ok = bcm.connectionPool[connStr]
	if ok {
		return conn, nil
	}

	newConn, err := bcm.openConnection(connStr)
	if err != nil {
		return nil, err
	}

	bcm.connectionPool[connStr] = newConn
	return newConn, nil
}

func (bcm *BasicConnectionManager) openConnection(connectionString string) (*sql.DB, error) {
	log.Printf("Connecting to %s", connectionString)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Println("Cannot connect: ", err.Error())
		return nil, err
	}

	log.Printf("Connected, Current open connections: %d", db.Stats().OpenConnections)

	return db, nil
}

func pingServer(db *sql.DB) error {
	log.Printf("Sending ping to SQL Server ...")
	err := db.Ping()
	if err != nil {
		log.Printf("Failed to sending ping due to: %s", err.Error())
		return err
	}
	log.Println("Ping succeeded")

	return nil
}
