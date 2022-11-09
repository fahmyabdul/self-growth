package transactions

import (
	"context"
	"errors"
	"net/http"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/app"
	"go.mongodb.org/mongo-driver/bson"
)

// Delete :
func (p *Transactions) Delete(transactionHash string) (interface{}, int, error) {
	mongoConn := app.Properties.Databases.MongoConn

	coll := mongoConn.Collection(p.CollectionName())

	result, err := coll.DeleteOne(context.Background(), bson.M{"transaction_hash": transactionHash})
	if err != nil {
		golibs.Log.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	if result.DeletedCount == 0 {
		golibs.Log.Println("| Transactions | Delete | ERROR | Requested transaction_hash not found")
		return nil, http.StatusNotFound, errors.New("requested transaction_hash not found")
	}

	return nil, http.StatusOK, nil
}
