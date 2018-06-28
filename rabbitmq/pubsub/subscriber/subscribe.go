package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/couchbase/gocb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// The SystemEvent represents an event
type SystemEvent struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	EventTime time.Time `json:"eventTime"`
}

func main() {
	// Echo instance
	e := echo.New()

	bucket := openBucket("pub-sub-demo-store")

	defer bucket.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/system-event/:id", getSystemEvent(bucket, "id"))
	e.POST("/system-event", createOrUpdateSystemEvent(bucket))
	e.PUT("/system-event", createOrUpdateSystemEvent(bucket))
	e.DELETE("/system-event/:id", deleteSystemEvent(bucket, "id"))

	// Start server
	e.Logger.Fatal(e.Start(":8923"))
}

func getSystemEvent(bucket *gocb.Bucket, parameterName string) func(c echo.Context) error {

	return func(c echo.Context) error {

		// Get parameter from URI (:id) case sensitive
		// If the parameter was not found,
		key := c.Param(parameterName)

		var doc *SystemEvent

		_, errGet := bucket.Get(key, &doc)

		if errGet != nil {
			errorMsg := fmt.Sprintf("Failed to get doc by [%s] due to %v.", key, errGet)
			log.Print(errorMsg)
			return c.JSON(http.StatusNotFound, errorMsg)
		}

		return c.JSON(http.StatusOK, doc)
	}
}

func deleteSystemEvent(bucket *gocb.Bucket, parameterName string) func(c echo.Context) error {
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

func createOrUpdateSystemEvent(bucket *gocb.Bucket) func(c echo.Context) error {

	return func(c echo.Context) error {
		doc := new(SystemEvent)
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

func openBucket(name string) *gocb.Bucket {

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

	log.Printf("Opening bucket %s.", name)

	bucket, err := cluster.OpenBucket(name, "")

	if err != nil {
		log.Printf("Failed to open buket %s due to %v", name, err)
		os.Exit(-1)
	}

	log.Printf("Bucket %s opened successfully.", name)

	return bucket
}
