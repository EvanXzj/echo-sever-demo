package main

import (
	"github.com/evanxzj/echo-server/api"
	"github.com/evanxzj/echo-server/models"
)

func main() {

	db := models.DB{}
	sv := api.Server{Db: &db}

	// setup
	db.Start()
	sv.Start()

	defer db.Client.Close()

}
