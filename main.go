package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	handler "github.com/moneybin/moneybin-paz/handlers"
	"github.com/moneybin/moneybin-paz/models"
)

func main() {

	models.InitDB()

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8881", nil))
}
