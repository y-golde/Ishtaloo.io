package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// somtehing about e
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		fmt.Printf("wowowowo");
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}