package mdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MDB encapsulates mongo.Databases
type MDB struct {
	client *mongo.Client
	rowdb  *mongo.Database // row database
	coldb  *mongo.Database // col database (storing columnar data)
}

// RowCollection returns the collection in the row database.
func (mdb *MDB) RowCollection(collectionName string) *mongo.Collection {
	return mdb.rowdb.Collection(collectionName)
}

// ColCollection returns the collection in the col database.
func (mdb *MDB) ColCollection(collectionName string) *mongo.Collection {
	return mdb.coldb.Collection(collectionName)
}

// NewMDB creates a new instance of MDB
func NewMDB(uri string, rowDBName string, colDBName string) (*MDB, error) {
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

	rowDB := client.Database(rowDBName)
	colDB := client.Database(colDBName)

	return &MDB{
		client: client,
		rowdb:  rowDB,
		coldb:  colDB,
	}, nil
}

// Disconnect disconnects from MongoDB
func (mdb *MDB) Disconnect(ctx context.Context) error {
	return mdb.client.Disconnect(ctx)
}

// Drop drops the databases
func (mdb *MDB) Drop(ctx context.Context) error {
	err := mdb.rowdb.Drop(ctx)
	if err != nil {
		return err
	}

	err = mdb.coldb.Drop(ctx)
	if err != nil {
		return err
	}

  return nil
}

// Ping pings MongoDB
func (mdb *MDB) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mdb.client.Ping(ctx, nil)
}

// InsertOneRow inserts a single record into the row database
func (mdb *MDB) InsertOneRow(ctx context.Context, collectionName string, document interface{}) error {
	collection := mdb.RowCollection(collectionName)
	_, err := collection.InsertOne(ctx, document)
	return err
}

// InsertOneCol inserts a single record into the col database
func (mdb *MDB) InsertOneCol(ctx context.Context, collectionName string, document interface{}) error {
	collection := mdb.ColCollection(collectionName)
	_, err := collection.InsertOne(ctx, document)
	return err
}
