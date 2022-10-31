package billing_event_worker

import (
	"github.com/anshumusaddi/billing_application/datastore"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/anshumusaddi/billing_application/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goccy/go-json"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkerPool struct {
	WorkerCount int
	store       *datastore.BillingApplicationDBStore
}

func InitEventWorkers(workerCount int, store *datastore.BillingApplicationDBStore) {
	pool := WorkerPool{
		workerCount,
		store,
	}
	for i := 1; i <= workerCount; i++ {
		go pool.InitWorkers(i)
	}
}

func (pool *WorkerPool) InitWorkers(workerID int) {
	consumer, err := datastore.GetKafkaConsumer(datastore.MessagingEventConsumerGroupID, datastore.MessagingEventTopic)
	if err != nil {
		logger.Error("kafka consumer failed to initialise : ", err.Error())
		return
	}
	logger.Info("worker created for ID: " + string(rune(workerID)))
	run := viper.GetBool("MESSAGING_EVENT_WORKER.Poll")
	for run == true {
		ev := consumer.Poll(1)
		switch e := ev.(type) {
		case *kafka.Message:
			message := e.Value
			messageEvent := models.MessageEvent{}
			err := json.Unmarshal(message, &messageEvent)
			if err != nil {
				logger.Error("kafka consumer failed to unmarshall data : ", err.Error())
			}
			logger.Debug("pushing the following event to db : ", messageEvent)
			err = pool.store.CreateOne(datastore.MessagingEventCollection, messageEvent)
			if err != nil {
				if mongo.IsDuplicateKeyError(err) {
					logger.Error("unique constrain violates for messaging_event collection")
				} else {
					logger.Error("error persisting to database, err: %s", err.Error())
				}
			}
		case kafka.Error:
			logger.Error("kafka consumer failed to fetch data : ", err.Error())
			run = false
		default:
			logger.Debug("kafka consumer poll ignored")
		}
	}
	return
}
