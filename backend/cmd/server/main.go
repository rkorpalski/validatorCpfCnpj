package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rkorpalski/validatorCpfCnpj/backend/api/handlers"
	"github.com/rkorpalski/validatorCpfCnpj/backend/api/routes"
	"github.com/rkorpalski/validatorCpfCnpj/backend/conf/env"
	"github.com/rkorpalski/validatorCpfCnpj/backend/pkg/cpfCnpj"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	db, err := initializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(handlers.CorsMiddleware())

	mongoRepository := cpfCnpj.NewMongoRepo(db)
	CnfCnpjService := cpfCnpj.NewCpfCnpjService(mongoRepository)

	validadorRoute := routes.NewValidatorRoute(CnfCnpjService)
	mainRouter := r.Group(env.ApiContext())
	validadorRoute.BuildRoutes(mainRouter)


	err = r.Run(env.AppBaseUrl())
	if err != nil {
		log.Fatal(err)
	}
}

func initializeDatabase() (*mongo.Database, error){
	client, err := mongo.NewClient(options.Client().ApplyURI(env.MongoBaseUrl()))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client.Database(env.AppDatabase()), nil
}