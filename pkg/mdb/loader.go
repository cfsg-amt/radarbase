package mdb

import (
	"context"
  "strconv"
  "strings"
  "fmt"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func (db *MDB) RowLoadToDB(data []map[string]interface{}, collectionName string) error {
	// Convert []map[string]interface{} to []interface{} for the MongoDB driver
	insertData := make([]interface{}, len(data))
	for i, v := range data {
		insertData[i] = v
	}

	// Get collection for data
	dataCollection := db.RowCollection(collectionName)

  // Update data into MongoDB
  for _, item := range insertData {
    itemMap, ok := item.(map[string]interface{})
    if !ok {
      return fmt.Errorf("item is not a map[string]interface{}, it's a %T", item)
    }
    idValue, ok := itemMap["_id"]
    if !ok {
      return fmt.Errorf("missing _id key in item map")
    }

    update := bson.M{
      "$set": item, // If the fields in item do not exist in the document, $set will add the fields.
    }

    filter := bson.M{"_id": idValue} // update the document with the same _id

    _, err := dataCollection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))

    if err != nil {
      return fmt.Errorf("could not update data into MongoDB: %w", err)
    }
  }

  fmt.Println("Row data updated into MongoDB! (", collectionName, ")")

	return nil
}

func (db *MDB) ColLoadToDB(data map[string][]interface{}, headers []string, collectionName string) error {
	// Create map to hold min and max data
	minData := make(map[string]float64)
	maxData := make(map[string]float64)

	// Initialize the min and max data with the first element of each header
	for _, header := range headers {
		if len(data[header]) > 0 {
			firstValue, ok := data[header][0].(float64)
			if ok {
				minData[header] = firstValue
				maxData[header] = firstValue
			} else {
        minData[header] = 0
        maxData[header] = 0
      }
		}
	}

	// Iterate through the data to find the min and max for each header
	for _, header := range headers {
		for _, value := range data[header] {
			floatValue, ok := value.(float64)
			if ok {
				if floatValue < minData[header] {
					minData[header] = floatValue
				}
				if floatValue > maxData[header] {
					maxData[header] = floatValue
				}
			} else {
        minData[header] = 0
        maxData[header] = 0
        continue;
      }
		}
	}

  // Create map to hold grouped data
	groupedData := make(map[int]map[string]interface{})
	for i := 0; i < 10; i++ {
		groupedData[i] = make(map[string]interface{})
		for _, header := range headers {
			groupedData[i][header] = []interface{}{}
		}
	}

  // Group data based on "時富雷達 (CR)" score
  for i, value := range data["時富雷達 (CR)"] {
    var scoreValue float64
    var err error

    switch v := value.(type) {
    case float64:
      scoreValue = v
    case string:
      trimmedV := strings.TrimSpace(v) // Trim spaces
      scoreValue, err = strconv.ParseFloat(trimmedV, 64)
      if err != nil {
        fmt.Printf("Value at index %d could not be parsed to float64: %v\n", i, err)
        continue
      }
    default:
      fmt.Printf("Value at index %d is not a float64 nor string, it's a %T\n", i, value)
      fmt.Println()
      continue
    }

    score := int(scoreValue);
    // Ensure score is within 1-10
    if score >= 0 && score < 10 {
      for _, header := range headers {
        groupedData[score][header] = append(groupedData[score][header].([]interface{}), data[header][i])
      }
    } else {
      return fmt.Errorf("score out of range: %d", score)
    }
  }

	// Get collection for headers
	headersCollection := db.ColCollection("headers")

	// Update headers into MongoDB headers collection with collection name as identifier
	headerDoc := bson.M{"_id": collectionName, "headers": headers}
	_, err := headersCollection.UpdateOne(context.Background(), bson.M{"_id": collectionName}, bson.M{"$set": headerDoc}, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("could not update headers into MongoDB: %w", err)
	}

	// Get collection for data and insert data into MongoDB
	dataCollection := db.ColCollection(collectionName)

	// Insert min and max data into the database
	minMaxDoc := bson.M{"_id": collectionName + "_min_max", "min": minData, "max": maxData}
	_, err = dataCollection.UpdateOne(context.Background(), bson.M{"_id": minMaxDoc["_id"]}, bson.M{"$set": minMaxDoc}, options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("could not update min/max data into MongoDB: %w", err)
	}

	// Insert other data into the database
	for score, groupedValues := range groupedData {
		groupedValues["_id"] = collectionName + strconv.Itoa(score)

		_, err := dataCollection.UpdateOne(context.Background(), bson.M{"_id": groupedValues["_id"]}, bson.M{"$set": groupedValues}, options.Update().SetUpsert(true))
		if err != nil {
			return fmt.Errorf("could not update data into MongoDB: %w", err)
		}
	}

  fmt.Println("Column data updated into MongoDB! (", collectionName, ")")

	return nil
}
