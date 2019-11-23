package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	handler "github.com/moneybin/moneybin-paz/handlers"
	"github.com/moneybin/moneybin-paz/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "paz"
	password = "password"
	dbname   = "moneybin"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	models.InitDB(psqlInfo)

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8881", nil))
}
