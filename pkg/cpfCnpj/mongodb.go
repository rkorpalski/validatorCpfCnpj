package cpfCnpj

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rkorpalski/validatorCpfCnpj/pkg/messages"
	"github.com/rkorpalski/validatorCpfCnpj/pkg/util"
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
	now, err := util.FormatTime("02-01-2006 15:04", time.Now())
	if err != nil {
		return errors.Wrap(err, messages.TimeFormatError)
	}
	cpfCnpj.CreateDate = now
	_, err = r.Db.Collection("document").InsertOne(context.Background(), cpfCnpj)
	if err != nil{
		return errors.Wrap(err, messages.SaveDocumentError)
	}

	return nil
}

func (r *MongoRepo) GetDocuments(isBlacklist bool) ([]CpfCnpj, error) {
	cursor, err := r.Db.Collection("document").Find(context.Background(), bson.D{{"blacklist", isBlacklist}})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	results := make([]CpfCnpj, 0)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var cpfCnpj CpfCnpj
		err := cursor.Decode(&cpfCnpj)
		if err != nil {
			return nil, errors.Wrap(err, messages.DocumentFindError)
		}
		results = append(results, cpfCnpj)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, messages.DocumentFindError)
	}
	cursor.Close(context.Background())

	return results, nil
}

func (r *MongoRepo) MoveToBlacklist(documentId string) error {
	_, err := r.Db.Collection("document").UpdateOne(context.Background(), bson.D{{"_id", documentId}},
														bson.D{{"$set", bson.D{{"blacklist", true}}}})
	if err != nil {
		return errors.Wrap(err, messages.BlacklistDocumentError)
	}
	return nil
}

func (r *MongoRepo) DeleteDocument(documentId string) error {
	_, err := r.Db.Collection("document").DeleteOne(context.Background(),  bson.D{{"_id", documentId}})
	if err != nil {
		return errors.Wrap(err, messages.DeleteDocumentError)
	}
	return nil
}

func (r *MongoRepo) RemoveFromBlacklist(documentId string) error {
	_, err := r.Db.Collection("document").UpdateOne(context.Background(), bson.D{{"_id", documentId}},
		bson.D{{"$set", bson.D{{"blacklist", false}}}})
	if err != nil {
		return errors.Wrap(err, messages.BlacklistRemoveError)
	}
	return nil
}

func (r *MongoRepo) FindByDocument(document string) ([]CpfCnpj, error) {
	var cpfCnpj CpfCnpj
	results := make([]CpfCnpj, 0)
	err := r.Db.Collection("document").FindOne(context.Background(), bson.D{{"number", document}}).Decode(&cpfCnpj)
	if err != nil {
		if err.Error() == messages.NoResultsMongoError {
			return results, nil
		}
		return nil, errors.Wrap(err, messages.FindDocumentError)
	}
	results = append(results, cpfCnpj)
	return results, nil
}