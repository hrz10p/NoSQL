package main

import (
	"context"
	"log"
	"main/pkg/logger"
	"main/pkg/services"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	logger.GetLogger().Info("Connected to MongoDB!")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

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
