package datastore

import (
	"context"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/goccy/go-json"
	"github.com/spf13/viper"
)

var topicPartMap = map[string]*kafka.TopicPartition{}

const (
	MessagingEventConsumerGroupID = "messaging_event_consumer"
	MessagingEventTopic           = "messaging_event"
)

func GetKafkaProducerConfigMap() kafka.ConfigMap {
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA.Host"),
	}
	return kafkaConfig
}

func GetKafkaConsumerConfigMap(groupID string) kafka.ConfigMap {
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": viper.GetString("KAFKA.Host"),
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}
	return kafkaConfig
}

func GetTopicPartition(topic string) *kafka.TopicPartition {
	topicPart, ok := topicPartMap[topic]
	if ok == false {
		topicPart = &kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		}
		topicPartMap[topic] = topicPart
	}
	return topicPart
}

func GetKafkaProducer() (*kafka.Producer, error) {
	config := GetKafkaProducerConfigMap()
	producer, err := kafka.NewProducer(&config)
	if err != nil {
		logger.Error("Unable to connect to kafka. err:", err.Error())
		return nil, err
	}
	err = createMessagingEventTopics()
	if err != nil {
		logger.Error("Unable to create topic in kafka. err:", err.Error())
		return nil, err
	}
	return producer, nil
}

func (store *BillingApplicationKafkaStore) ProduceEvent(topic string, message interface{}) error {
	topicPart := GetTopicPartition(topic)
	rawMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = store.p.Produce(&kafka.Message{
		TopicPartition: *topicPart,
		Value:          rawMessage,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

type BillingApplicationKafkaStore struct {
	p *kafka.Producer
}

func NewBillingApplicationKafkaStore(p *kafka.Producer) *BillingApplicationKafkaStore {
	return &BillingApplicationKafkaStore{p: p}
}

func createMessagingEventTopics() error {
	config := GetKafkaProducerConfigMap()
	adminClient, err := kafka.NewAdminClient(&config)
	if err != nil {
		logger.Error("Unable to connect to kafka. err:", err.Error())
		return err
	}
	_, err = adminClient.CreateTopics(context.Background(), []kafka.TopicSpecification{{
		Topic:             MessagingEventTopic,
		NumPartitions:     1,
		ReplicationFactor: 1}},
	)
	if err != nil {
		logger.Error("unable to create topic in kafka. err:", err.Error())
		return err
	}
	return nil
}

func GetKafkaConsumer(groupID string, topic string) (*kafka.Consumer, error) {
	config := GetKafkaConsumerConfigMap(groupID)
	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		logger.Error("unable to connect to kafka. err:", err.Error())
		return nil, err
	}
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		logger.Error("unable to subscribe to kafka. err:", err.Error())
		return nil, err
	}
	return consumer, nil
}
