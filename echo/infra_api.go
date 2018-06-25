package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	// Echo instance
	e := echo.New()

	bucket := openBucket()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api", getAPIDoc(bucket))

	e.POST("/api", createAPIDoc(bucket))

	// Start server
	e.Logger.Fatal(e.Start(":8323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getAPIDoc(bucket *gocb.Bucket) func(c echo.Context) error {

	return func(c echo.Context) error {
		key := c.QueryParam("ID")

		var doc *APIDefinition

		_, errGet := bucket.Get(key, &doc)

		if errGet != nil {
			errorMsg := fmt.Sprintf("Failed to get doc by [%s] due to %v.", key, errGet)
			log.Print(errorMsg)
			return c.JSON(http.StatusInternalServerError, errorMsg)
		}

		if doc == nil {
			log.Printf("Doc not found for key %s.", key)
			return c.JSON(http.StatusNotFound, "API Doc not found.")
		}

		return c.JSON(http.StatusOK, doc)
	}
}

func createAPIDoc(bucket *gocb.Bucket) func(c echo.Context) error {

	return func(c echo.Context) error {
		doc := new(APIDefinition)
		if err := c.Bind(doc); err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprint(err))
		}

		cas, err := bucket.Upsert(doc.ID, doc, 0)
		if err == nil {
			return c.JSON(http.StatusCreated, fmt.Sprintf("CAS: %v Key: %s", cas, doc.ID))
		}

		return c.String(http.StatusInternalServerError, fmt.Sprint(err))
	}
}

func openBucket() *gocb.Bucket {
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

	return bucket
}
