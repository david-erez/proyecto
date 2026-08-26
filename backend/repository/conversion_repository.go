package repository

import (
	"context"

	"backend/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (r *ConversionRepository) UpdateByID(ctx context.Context, id string, update bson.M) (models.ConversionRecord, error) {
	oid, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return models.ConversionRecord{}, err

	}

	filter := bson.M{"_id": oid}
	updateDoc := bson.M{"$set": update}

	result := r.collection.FindOneAndUpdate(
		ctx,
		filter,
		updateDoc,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	var record models.ConversionRecord
	if err := result.Decode(&record); err != nil {
		return models.ConversionRecord{}, err

	}

	return record, nil

}

func (r *ConversionRepository) DeleteByID(ctx context.Context, id string) error {
	oid, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
