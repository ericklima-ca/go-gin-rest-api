package main

import (
	"github.com/ericklima-ca/go-gin-rest-api/database"
	"github.com/ericklima-ca/go-gin-rest-api/routes"
)

func main() {
	database.Connect()
	router := routes.Handler()
	router.Run(":8888")
}
