package DB

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func InitClient() (error) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	fmt.Printf("DB URL: " + os.Getenv("DB_URL"))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		log.Fatal("cannot connect to mongo server" + err.Error())
		return err
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	return nil
}