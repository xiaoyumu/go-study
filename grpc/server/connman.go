package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	rda "github.com/xiaoyumu/go-study/grpc/proto"
)

// ConnectionManager contains function definitions related to a connection pooling management
type ConnectionManager interface {
	// GetConnection returns a ptr to sql.DB based on given connection string
	GetConnection(connStr string) (*sql.DB, error)
	// BuildConnectionString builds connection string based on the DBRequest
	BuildConnectionString(req *rda.DBRequest) (string, error)
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

	// Check if a connection already exists before lock it
	conn, ok := bcm.connectionPool[connStr]
	if ok {
		// Reuse it if already exists
		return conn, nil
	}

	// Now try to lock it for double check
	bcm.mutex.Lock()
	defer bcm.mutex.Unlock() // Release the lock after function returns

	// Check again if the connection already exists in the pool
	conn, ok = bcm.connectionPool[connStr]
	if ok {
		// Reuse it if already exists
		return conn, nil
	}

	// Create a new one
	newConn, err := bcm.openConnection(connStr)
	if err != nil {
		return nil, err
	}

	// Add the new connection to the pool
	bcm.connectionPool[connStr] = newConn
	return newConn, nil
}

// BuildConnectionString builds connection string based on the DBRequest
func (bcm *BasicConnectionManager) BuildConnectionString(req *rda.DBRequest) (string, error) {
	// Sample Connection string:
	// sqlserver://dev:d3v@192.168.1.154:1433?database=godemo&connection+timeout=30
	conn := fmt.Sprintf("sqlserver://%s:%s@%s:%v?database=%s&connection+timeout=30",
		req.ServerInfo.UserId,
		req.ServerInfo.Password,
		req.ServerInfo.Server,
		req.ServerInfo.Port,
		req.ServerInfo.Database)
	return conn, nil
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
