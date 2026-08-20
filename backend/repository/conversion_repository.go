package repository

import (
	"context"

	"backend/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ConversionRepository struct {
	collection *mongo.Collection
}

//crear el repositorio en base de un cliente de mongo ya conectado

func NewConversionRepository(client *mongo.Client, dbNmae string) *ConversionRepository {
	collection := client.Database(dbNmae).Collection("conversions")
	return &ConversionRepository{collection: collection}
}

func (r *ConversionRepository) Save(ctx context.Context, record models.ConversionRecord) (models.ConversionRecord, error) {
	result, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		return models.ConversionRecord{}, err
	}
	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
		record.ID = oid.Hex()
	}
	return record, nil
}

func (r *ConversionRepository) FindAll(ctx context.Context) ([]models.ConversionRecord, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	records := []models.ConversionRecord{}
	if err := cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *ConversionRepository) FindByID(ctx context.Context, id string) (models.ConversionRecord, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return models.ConversionRecord{}, err
	}

	var record models.ConversionRecord
	err = r.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&record)
	if err != nil {
		return models.ConversionRecord{}, err
	}

	return record, nil
}
