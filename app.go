package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gsvaldevieso/go-dream-architecture/infrastructure/database"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	_ = database.Connect()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
