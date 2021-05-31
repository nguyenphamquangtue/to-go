package main

import (
	"log"
	"togo/api"
	"togo/app"
	"togo/model"
)

const port = ":5050"

func main() {
	app.Initialize()
	model.Initialize()
	defer model.ClosePGConnection()

	router := api.NewRouter()
	log.Fatal(router.Run(port))
}
