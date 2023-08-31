package main

import (
	"min/api"
	"min/model"
	"min/service"
)

func main() {
	api.StartCron()
	model.Init()

	//serve := service.NewRouter()

	service.Run(":3000")
}
