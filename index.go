package main

import (
	"log"
	"net/http"
	"ngnt_marketplace/controllers"
)

func main() {
	mux := http.NewServeMux()
	port := ":6060"
	mux.HandleFunc("/", controllers.HealthHandler)
	mux.HandleFunc("/order", controllers.CreateOrderHandler)
	log.Println("Starting server on " + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
