package transactions

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/fahmyabdul/golibs"
	"github.com/fahmyabdul/self-growth/data-app/app"
)

func (p *Transactions) Create() (interface{}, int, error) {
	mongoConn := app.Properties.Databases.MongoConn

	coll := mongoConn.Collection(p.CollectionName())

	p.CreatedAt, _ = strconv.Atoi(time.Now().Format("20060102150405"))
	p.UpdatedAt, _ = strconv.Atoi(time.Now().Format("20060102150405"))

	_, err := coll.InsertOne(context.Background(), p)
	if err != nil {
		golibs.Log.Printf(err.Error())
		return p, http.StatusInternalServerError, err
	}

	return p, 0, nil
}
