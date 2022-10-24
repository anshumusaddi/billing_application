package main

import (
	"github.com/anshumusaddi/billing_application/config"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/anshumusaddi/billing_application/models"
	"github.com/anshumusaddi/billing_application/routes"
	"github.com/spf13/viper"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		return
	}

	logger.InitLogger(viper.GetString("LOG.Level"))
	defer logger.Sync()

	var events []models.MessageEvent
	engine := routes.InitRoutes(&events)
	err = engine.Run(":8080")
	if err != nil {
		return
	}
}
