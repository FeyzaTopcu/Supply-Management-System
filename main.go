package main

import (
	"net/http"

	hand "./handler"
	"github.com/gorilla/mux"

	"log"
	
)

var router = mux.NewRouter()

func main() {
	log.Println("Server Starting")
	router.HandleFunc("/register/users", hand.Handler)
	router.HandleFunc("/signup", hand.SignUpHandler)
	router.HandleFunc("/login", hand.LoginHandler)

	router.HandleFunc("/api/products", hand.GetProductsHandler).Methods("GET")
	router.HandleFunc("/api/products/{id}", hand.GetProductHandler).Methods("GET")
	router.HandleFunc("/api/product", hand.PostProductHandler).Methods("POST")
	router.HandleFunc("/api/products/{id}", hand.PutProductHandler).Methods("Put")
	router.HandleFunc("/api/products/{id}", hand.DeleteProductHandler).Methods("Delete")
	
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
	log.Println("Server ending...")
}
