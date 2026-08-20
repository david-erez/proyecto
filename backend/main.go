package main

import (
	"log"
	"net/http"

	"backend/config"
	"backend/db"
	"backend/handlers"
	"backend/repository"
	"backend/services"
)

func main() {
	cfg := config.Load()

	client := db.Connect(cfg.MongoURI)

	repo := repository.NewConversionRepository(client, cfg.DatabaseName)
	service := services.NewConversionService(repo)
	handler := handlers.NewConversionHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/upload", handler.Upload)
	mux.HandleFunc("/api/conversions", handler.List)

	log.Println("Backend ejecutándose en el puerto " + cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, withCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
