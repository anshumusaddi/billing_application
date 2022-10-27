package main

import (
	"github.com/anshumusaddi/billing_application/config"
	"github.com/anshumusaddi/billing_application/datastore"
	"github.com/anshumusaddi/billing_application/logger"
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

	if err := datastore.InitDB(); err != nil {
		logger.Fatal("unable to connect to DB, err:", err.Error())
	}

	database := datastore.GetDb()
	store := datastore.NewBillingApplicationDBStore(database)

	engine := routes.InitRoutes(store)
	err = engine.Run(":8080")
	if err != nil {
		return
	}
}
