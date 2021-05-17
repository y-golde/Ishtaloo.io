package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rootController "ishtaloo.io/API"
	"ishtaloo.io/DB"
	entities "ishtaloo.io/Entities"
)

var sseChannel entities.SSEChannel

func main() {
	DB.InitClient()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// change with dotenv later to remove origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowCredentials: true,
	}))

	// Route => handler
	rootController.RootController(e, &sseChannel)

	// initialize global variable
	sseChannel = entities.SSEChannel{
		Clients:  make([]chan string, 0),
		Notifier: make(chan string),
	}

	// done signal to go routine.
	done := make(chan interface{})
	defer close(done)

	// run our broadcaster go routine.
	go broadcaster(done, &sseChannel)

	e.GET("/sse", sseRequest)
	e.GET("/log", func(c echo.Context) error { return logRequest(c, &sseChannel) })

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func broadcaster(done <-chan interface{}, sseChanel *entities.SSEChannel) {
	fmt.Println("Broadcaster Started.")
	for {
		select {
		case <-done:
			return
		case data := <-sseChannel.Notifier:
			for _, channel := range sseChannel.Clients {
				channel <- data
			}
		}
	}
}

func logRequest(c echo.Context, sseChanel *entities.SSEChannel) error {
	var body string
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusForbidden, "body bad format.")
	}

	sseChannel.Notifier <- "users changed"
	return nil
}

func sseRequest(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().WriteHeader(http.StatusOK)

	sseChan := make(chan string)
	sseChannel.Clients = append(sseChannel.Clients, sseChan)

	d := make(chan interface{})
	defer close(d)
	defer fmt.Println("Closing channel.")

	for {
		select {
		case <-d:
			close(sseChan)
			return c.NoContent(http.StatusOK)
		case data := <-sseChan:
			fmt.Fprintf(c.Response(), "new event: %v \n\n", data)
			c.Response().Flush()
		}
	}
}
