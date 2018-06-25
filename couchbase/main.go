package main

import (
	"fmt"
	"log"
	"os"

	"github.com/couchbase/gocb"
)

// The APIDefinition define the root object of an API in the infrastucture
type APIDefinition struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Hosts       []Host   `json:"hosts"`
	URIs        []string `json:"uris"`
}

// The Host defines the basic information of a Host
type Host struct {
	Address string `json:"address"`
	Port    int32  `json:"port"`
}

func main() {

	const database string = "infrastructure"
	const connectionString string = "couchbase://192.168.1.51,192.168.1.52,192.168.1.53"

	log.Printf("Connecting to %s.", connectionString)

	cluster, connectionError := gocb.Connect(connectionString)

	if connectionError != nil {
		log.Printf("Failed to connect to %s due to %v", connectionString, connectionError)
		os.Exit(-1)
	}

	log.Printf("Connected successfully.")

	authError := cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "dev",
		Password: "p@ssw0rd",
	})

	if authError != nil {
		log.Printf("Authentication failed.")
		os.Exit(-1)
	}

	log.Printf("Opening bucket %s.", database)

	bucket, err := cluster.OpenBucket(database, "")

	if err != nil {
		log.Printf("Failed to open buket %s due to %v", database, err)
		os.Exit(-1)
	}

	log.Printf("Bucket %s opened successfully.", database)

	key := "api-10001"

	var doc *APIDefinition

	_, errGet := bucket.Get(key, &doc)

	if errGet != nil {
		log.Printf("Failed to get doc by [%s] due to %v.", key, errGet)
		os.Exit(-1)
	}

	if doc == nil {
		log.Printf("New doc.")
	} else {
		log.Printf("Existing doc.")
		log.Println(*doc)
	}

	casUpsert, errUpsert := bucket.Upsert(
		key,
		APIDefinition{
			ID:          key,
			Name:        "demo API",
			Description: "The demo api to demostrate on how to use couchbase.",
			Version:     "v1",
			Hosts: []Host{
				Host{Address: "127.0.0.1", Port: 8080},
				Host{Address: "127.0.0.1", Port: 8081},
			},
			URIs: []string{
				"/api/v1/demo-api",
				"/api/v1/demo-api/test",
			},
		}, 0)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to upsert a doc due to %v\n", errUpsert)
		os.Exit(-1)
	}

	fmt.Println(casUpsert)
}
