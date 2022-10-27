package datastore

import (
	"context"
	"fmt"
	"github.com/anshumusaddi/billing_application/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	MessagingEventCollection = "messaging_event"
)

var db *mongo.Database

func GetDb() *mongo.Database {
	return db
}

func getMongoDbUri() string {
	mongodbUri := fmt.Sprintf("mongodb://%s:%s",
		viper.GetString("DB.Host"), viper.GetString("DB.Port"))
	return mongodbUri
}

func InitDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(getMongoDbUri())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("Unable to connect to DB. err:", err.Error())
		return err
	}
	if err = client.Ping(ctx, nil); err != nil {
		logger.Error("Unable to ping DB. err:", err.Error())
		return err
	}

	db = client.Database(viper.GetString("DB.Name"))

	err = createMessagingEvents(db)
	if err != nil {
		return err
	}

	return nil
}

type BillingApplicationDBStore struct {
	db *mongo.Database
}

func NewBillingApplicationDBStore(db *mongo.Database) *BillingApplicationDBStore {
	return &BillingApplicationDBStore{db: db}
}

func (store *BillingApplicationDBStore) GetDb() *mongo.Database {
	return store.db
}

func (store *BillingApplicationDBStore) CreateOne(collection string, document interface{}) error {
	_, err := store.db.Collection(collection).InsertOne(context.Background(), document)
	return err
}

func (store *BillingApplicationDBStore) RemoveDeletedDocuments(query interface{}) interface{} {
	qM, ok := query.(bson.M)
	if ok {
		qM["deleted_at"] = bson.M{"$exists": false}
		return qM
	}
	qD, ok := query.(bson.D)
	if ok {
		qD = append(qD, bson.E{Key: "deleted_at", Value: bson.M{"$exists": false}})
		return qD
	}
	return query
}

func (store *BillingApplicationDBStore) Find(collection string, document interface{}, query bson.M) error {
	c, err := store.db.Collection(collection).Find(context.Background(), store.RemoveDeletedDocuments(query))
	if err != nil {
		return err
	}
	err = c.All(context.Background(), document)
	if err != nil {
		return err
	}

	return nil
}

func createMessagingEvents(db *mongo.Database) error {
	collection := db.Collection(MessagingEventCollection)
	indexModel := []mongo.IndexModel{
		{
			Keys: bson.D{primitive.E{Key: "time", Value: 1}, primitive.E{Key: "customer_id", Value: 1},
				primitive.E{Key: "deleted_at", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("customer_time_unique_index"),
		},
	}
	_, err := collection.Indexes().CreateMany(context.Background(), indexModel)
	if err != nil {
		logger.Error("not able to create index on %s collection. err: %s", MessagingEventCollection, err.Error())
		return err
	}
	return nil
}
