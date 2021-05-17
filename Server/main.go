package main

import (
	"fmt"
	"log"
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

	// initialize global variable
	sseChannel = entities.SSEChannel{
		Clients:  make([]chan string, 0),
		Notifier: make(chan string),
	}

	// Route => handler
	rootController.RootController(e, &sseChannel)

	// done signal to go routine.
	done := make(chan interface{})
	defer close(done)

	// run our broadcaster go routine.
	go broadcaster(done, &sseChannel)

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
