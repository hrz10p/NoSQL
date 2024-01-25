package main

import (
	"context"
	"fmt"
	"log"
	"main/pkg/logger"
	"main/pkg/services"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	Service *services.Service
	Router  *http.ServeMux
	Logger  *logger.Logger
}

func NewApplication() *Application {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router := http.NewServeMux()

	return &Application{
		Service: services.NewService(client.Database("nsqlproj")),
		Router:  router,
		Logger:  logger.GetLogger(),
	}
}

func (app *Application) Start(addr string) error {
	app.Logger.Info("starting server... on localhost:8080")
	app.InitializeRoutes()
	return http.ListenAndServe(addr, app.Router)
}
