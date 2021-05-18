package main

import (
	"log"
	"net/http"

	"github.com/obbap1/ngnt_marketplace/controllers"
)

func main() {
	http.HandleFunc("/", controllers.HealthHandler)
	log.Fatal(http.ListenAndServe(":6060", nil))
}
