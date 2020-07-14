package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	adapter "github.com/gsvaldevieso/go-dream-architecture/adapter/database"
	"github.com/gsvaldevieso/go-dream-architecture/domain/entity"
	"github.com/gsvaldevieso/go-dream-architecture/infrastructure/database"
	"github.com/gsvaldevieso/go-dream-architecture/repository"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	mysql, _ := database.NewMySQL()
	dbAdapter := adapter.NewRelational(mysql.DB)
	orderRepo := repository.NewOrderRepo(dbAdapter)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var response = struct {
			Status int `json:"status"`
		}{
			Status: http.StatusOK,
		}

		err := orderRepo.Store(entity.Order{ID: "ID"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response.Status = http.StatusInternalServerError
			_ = json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(response)
	})

	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
