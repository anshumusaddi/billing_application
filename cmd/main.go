package main

import (
	"github.com/anshumusaddi/billing_application/billing_event_worker"
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

	producer, err := datastore.GetKafkaProducer()
	if err != nil {
		return
	}
	kafkaStore := datastore.NewBillingApplicationKafkaStore(producer)

	billing_event_worker.InitEventWorkers(viper.GetInt("MESSAGING_EVENT_WORKER.Count"), store)

	engine := routes.InitRoutes(store, kafkaStore)
	err = engine.Run(":8080")
	if err != nil {
		return
	}
}
