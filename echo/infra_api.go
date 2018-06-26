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

	defer bucket.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/:id", getAPIDoc(bucket, "id"))
	e.POST("/api", createOrUpdateAPIDoc(bucket))
	e.PUT("/api", createOrUpdateAPIDoc(bucket))
	e.DELETE("/api/:id", deleteAPIDoc(bucket, "id"))

	// Start server
	e.Logger.Fatal(e.Start(":8323"))
}

func deleteAPIDoc(bucket *gocb.Bucket, parameterName string) func(c echo.Context) error {
	return func(c echo.Context) error {
		key := c.Param(parameterName)

		var cas gocb.Cas

		_, errGet := bucket.Remove(key, cas)

		if errGet != nil {
			errorMsg := fmt.Sprintf("Failed to get The doc [ID: %s] due to %v.", key, errGet)
			log.Print(errorMsg)
			return c.JSON(http.StatusNotFound, errorMsg)
		}

		return c.JSON(http.StatusOK, "Removed")
	}
}

func getAPIDoc(bucket *gocb.Bucket, parameterName string) func(c echo.Context) error {

	return func(c echo.Context) error {

		// Get parameter from URI (:id) case sensitive
		// If the parameter was not found,
		key := c.Param(parameterName)

		var doc *APIDefinition

		_, errGet := bucket.Get(key, &doc)

		if errGet != nil {
			errorMsg := fmt.Sprintf("Failed to get doc by [%s] due to %v.", key, errGet)
			log.Print(errorMsg)
			return c.JSON(http.StatusNotFound, errorMsg)
		}

		return c.JSON(http.StatusOK, doc)
	}
}

func createOrUpdateAPIDoc(bucket *gocb.Bucket) func(c echo.Context) error {

	return func(c echo.Context) error {
		doc := new(APIDefinition)
		// Bind the request body to the doc entity instance
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
