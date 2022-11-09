package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *Transactions) GetAll() (interface{}, int, int, error) {
	mongoConn := app.Properties.Databases.MongoConn

	coll := mongoConn.Collection(p.CollectionName())

	var result []Transactions
	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		golibs.Log.Println(err.Error())
		return nil, 0, http.StatusInternalServerError, err
	}

	// iterate over docs using Next()
	for cursor.Next(context.Background()) {
		// declare a result BSON object
		var resultCursor Transactions
		err := cursor.Decode(&resultCursor)
		if err != nil {
			continue
		}

		result = append(result, resultCursor)
	}

	if len(result) == 0 {
		golibs.Log.Println("| Transactions | GetByUsername | ERROR | Collection is empty")
		return nil, 0, http.StatusNotFound, fmt.Errorf("collection is empty")
	}

	return &result, len(result), http.StatusOK, nil
}

// GetByUsername :
func (p *Transactions) GetByUsername(username string) (interface{}, int, error) {
	mongoConn := app.Properties.Databases.MongoConn

	coll := mongoConn.Collection(p.CollectionName())

	var result Transactions
	err := coll.FindOne(context.TODO(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		golibs.Log.Println("| Transactions | GetByUsername | ERROR | No document found with given username")
		return nil, http.StatusNotFound, fmt.Errorf("no document found with given username")
	}

	if result == (Transactions{}) {
		golibs.Log.Println("| Transactions | GetByUsername | ERROR | No document found with given username")
		return nil, http.StatusNotFound, fmt.Errorf("no document found with given username")
	}

	return &result, http.StatusOK, nil
}

// GetByFilter :
func (p *Transactions) GetByFilter(jsonData []byte) (interface{}, int, error) {
	mongoConn := app.Properties.Databases.MongoConn

	coll := mongoConn.Collection(p.CollectionName())

	var requestBody map[string]interface{}
	err := json.Unmarshal(jsonData, &requestBody)
	if err != nil {
		golibs.Log.Println("| Transactions | GetByFilter | ERROR | Invalid json request")
		return nil, http.StatusBadRequest, fmt.Errorf("invalid json request")
	}

	filter := bson.M{}

	for k, v := range requestBody {
		filter[k] = v
	}

	if _, ok := filter["_id"]; ok {
		objectID, err := primitive.ObjectIDFromHex(filter["_id"].(string))
		if err != nil {
			return nil, http.StatusBadRequest, fmt.Errorf("invalid json request")
		}

		filter["_id"] = bson.M{
			"$eq": objectID,
		}
	}

	var result Transactions
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		golibs.Log.Println("| Transactions | GetByFilter | ERROR | No document found with given filter")
		return nil, http.StatusNotFound, fmt.Errorf("no document found with given filter")
	}

	if result == (Transactions{}) {
		golibs.Log.Println("| Transactions | GetByFilter | ERROR | No document found with given filter")
		return nil, http.StatusNotFound, fmt.Errorf("no document found with given filter")
	}

	return &result, http.StatusOK, nil
}
