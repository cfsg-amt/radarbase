package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Database encapsulates mongo.Database
type Database struct {
	client     *mongo.Client
	database   *mongo.Database
}

// Collection returns the collection in the database.
func (db *Database) Collection(collectionName string) *mongo.Collection {
	return db.database.Collection(collectionName)
}

// NewDatabase creates a new instance of Database
func NewDatabase(uri string, database string) (*Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(database)

	return &Database{
		client:     client,
		database:   db,
	}, nil
}

// Disconnect disconnects from MongoDB
func (db *Database) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

// Drop drops the database
func (db *Database) Drop(ctx context.Context) error {
	return db.database.Drop(ctx)
}

// Ping pings MongoDB
func (db *Database) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
	return db.client.Ping(ctx, nil)
}

// InsertOne inserts a single record into MongoDB
func (db *Database) InsertOne(ctx context.Context, collectionName string, document interface{}) error {
	collection := db.Collection(collectionName)
	_, err := collection.InsertOne(ctx, document)
	return err
}
