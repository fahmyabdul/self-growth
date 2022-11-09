package transactions

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p *Transactions) Update(transactionHash string) (interface{}, int, error) {
	mongoConn := app.Properties.Databases.MongoConn

	coll := mongoConn.Collection(p.CollectionName())

	var existingTransaction Transactions
	err := coll.FindOne(context.TODO(), bson.M{"transaction_hash": transactionHash}).Decode(&existingTransaction)
	if err != nil {
		golibs.Log.Println("| Transactions | Update | ERROR | Transaction not found")
		return nil, http.StatusBadRequest, fmt.Errorf("transaction not found")
	}

	p.UpdatedAt, _ = strconv.Atoi(time.Now().Format("20060102150405"))
	update := bson.M{
		"$set": p,
	}

	filter := bson.M{"transaction_hash": bson.M{"$eq": transactionHash}}

	updateOptions := (&options.FindOneAndUpdateOptions{}).SetReturnDocument(options.After)

	resultUpdate := coll.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		updateOptions,
	)

	if resultUpdate.Err() != nil {
		golibs.Log.Println("| Transactions | Update | ERROR |", resultUpdate.Err())
		return nil, http.StatusInternalServerError, resultUpdate.Err()
	}

	var result Transactions
	err = resultUpdate.Decode(&result)
	if err != nil {
		golibs.Log.Println("| Transactions | Update | ERROR |", err.Error())
		return nil, http.StatusInternalServerError, err
	}

	return &result, http.StatusOK, nil
}
