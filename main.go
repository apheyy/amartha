package main

import (
	"amarthaTest/config"
	"amarthaTest/controller"
	"log"
)

func main() {
	log.Println("Starting amartha case study")
	cfg := config.InitConfig()

	controller.HandleAPI(cfg)
}
