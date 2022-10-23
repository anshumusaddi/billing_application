package main

import (
	"github.com/anshumusaddi/billing_application/models"
	"github.com/anshumusaddi/billing_application/routes"
)

func main() {
	var events []models.MessageEvent
	engine := routes.InitRoutes(&events)
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
