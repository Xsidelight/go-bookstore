package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Xsidelight/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}

	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)

	port := os.Getenv("PORT")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:"+port, r))
}
