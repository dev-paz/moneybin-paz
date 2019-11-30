package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	handler "github.com/moneybin/moneybin-paz/handlers"
	"github.com/moneybin/moneybin-paz/models"
)

func main() {
	port := os.Getenv("PORT")

	models.InitDB()

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
