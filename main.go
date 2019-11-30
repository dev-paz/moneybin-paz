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
	user     = "qpivtmamwliml"
	password = "a3184750c3ee8bb244ecb21efba4a5ecbc0dc442df9f251f657012ead1041aa2"
	dbname   = "dcf8v4unmrcm0k"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	models.InitDB(psqlInfo)

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8881", nil))
}
