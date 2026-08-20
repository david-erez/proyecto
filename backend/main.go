package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	mongoURI := "mongodb://mongodb:27017"

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error conectando a MongoDB:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
	log.Fatal("MongoDB no responde:", err)
	}

	log.Println("Conexión con MongoDB establecida correctamente")

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		response := Response{
			Message: "Backend Go conectado correctamente con MongoDB",
		}

		json.NewEncoder(w).Encode(response)
	})
	log.Println("Backend ejecutándose en el puerto 8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
