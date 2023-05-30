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
	collection *mongo.Collection
}

// Collection returns the collection in the database.
func (db *Database) Collection() *mongo.Collection {
	return db.collection
}

// NewDatabase creates a new instance of Database
func NewDatabase(uri string, database string, collection string) (*Database, error) {
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
	coll := db.Collection(collection)

	return &Database{
		client:     client,
		database:   db,
		collection: coll,
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

// InsertOne inserts a single document into MongoDB
func (db *Database) InsertOne(ctx context.Context, document interface{}) error {
	_, err := db.collection.InsertOne(ctx, document)
	return err
}
