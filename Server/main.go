package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"ishtaloo.io/DB"
	"ishtaloo.io/DB/Collections"
)

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

	// somtehing about e
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		fmt.Printf("wowowowo");
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	
	e.GET("/hello", func(c echo.Context) error {
		usersCollection := Collections.GetUsersCollection()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		
		cur, _ := usersCollection.Find(ctx, bson.M{})

		var users []bson.M
		err = cur.All(ctx, &users)

		return c.JSON(200, users)
	})

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}