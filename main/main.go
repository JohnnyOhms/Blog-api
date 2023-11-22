package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JohnnyOhms/BlogPost-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BlogRoutes(r)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type:", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Go Server end point with sql")
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("server running on port 8080 and DB connected")
	}
}
