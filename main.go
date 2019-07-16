package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rkorpalski/validatorCpfCnpj/api/handlers"
	"github.com/rkorpalski/validatorCpfCnpj/api/routes"
	"github.com/rkorpalski/validatorCpfCnpj/pkg/cpfCnpj"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	os.Setenv("TZ", "America/Sao_Paulo")
	db, err := initializeDatabase()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	r := gin.Default()
	r.Use(handlers.CorsMiddleware())

	mongoRepository := cpfCnpj.NewMongoRepo(db)
	CnfCnpjService := cpfCnpj.NewCpfCnpjService(mongoRepository)

	validadorRoute := routes.NewValidatorRoute(CnfCnpjService)
	mainRouter := r.Group("/validator")
	validadorRoute.BuildRoutes(mainRouter)


	err = r.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func initializeDatabase() (*mongo.Database, error){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client.Database("CpfCnpj"), nil
}
