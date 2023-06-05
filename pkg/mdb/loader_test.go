package mdb_test

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestPrintAllItems(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.ColCollection("test").Find(ctx, bson.M{})
	if err != nil {
		t.Fatalf("failed to find documents: %v", err)
	}

	if cursor.RemainingBatchLength() == 0 {
		t.Fatalf("No documents found.")
	}

	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			t.Fatalf("cursor.Decode() error: %v", err)
		}
		t.Log(result)
	}

	if err := cursor.Err(); err != nil {
		t.Fatalf("cursor.Err(): %v", err)
	}

	if err := cursor.Close(ctx); err != nil {
		t.Fatalf("cursor.Close() error: %v", err)
	}
}
