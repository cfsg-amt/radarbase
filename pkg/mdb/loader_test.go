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

	cursor, err := db.Collection().Find(ctx, bson.M{})
	if err != nil {
		t.Fatalf("failed to find documents: %v", err)
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		t.Fatalf("failed to decode documents: %v", err)
	}

	for _, result := range results {
		t.Log(result)
	}
}
