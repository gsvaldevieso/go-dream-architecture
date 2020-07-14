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
		err := orderRepo.Store(entity.Order{ID: "ID"})
		if err != nil {
			response{http.StatusInternalServerError}.send(w)
			return
		}
		response{http.StatusCreated}.send(w)
	})

	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type response struct {
	Status int `json:"status"`
}

func (r response) send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r)
}
