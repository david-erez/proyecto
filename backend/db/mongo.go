package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(uri string) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Error conectando a mongodb", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("Error mongo db no responde", err)
	}

	log.Println("conexion a la base de datos establecida")

	return client
}
