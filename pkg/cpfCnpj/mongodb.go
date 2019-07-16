package cpfCnpj

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoRepo struct {
	Db *mongo.Database
}

func NewMongoRepo(db *mongo.Database) *MongoRepo{
	return &MongoRepo{
		Db: db,
	}
}

func (r *MongoRepo) Save(cpfCnpj CpfCnpj) error {
	cpfCnpj.Id = primitive.NewObjectID().Hex()
	fmt.Println(time.Now())
	cpfCnpj.CreateDate = time.Now()
	_, err := r.Db.Collection("document").InsertOne(context.Background(), cpfCnpj)
	if err != nil{
		return err
	}

	return nil
}

func (r *MongoRepo) GetAllDocuments() ([]CpfCnpj, error) {
	cursor, err := r.Db.Collection("document").Find(context.Background(), bson.D{{"blacklist", false}})
	if err != nil {
		return nil, err
	}

	results := make([]CpfCnpj, 0)
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var cpfCnpj CpfCnpj
		err := cursor.Decode(&cpfCnpj)
		if err != nil {
			return nil, err
		}
		cpfCnpj.CreateDate = cpfCnpj.CreateDate.In(loc)
		results = append(results, cpfCnpj)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(context.Background())

	return results, nil
}
