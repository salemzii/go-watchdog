package databases

import (
	"context"

	"github.com/salemzii/go-watchdog/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeMongodbQueryCheck(db *Database) service.ServiceCheck {

	uri, err := db.DSNMongoDb()
	if err != nil {
		service.HandleError("mongodb", err)
	}
	if db.Uri_Only() {
		uri = db.UriOnly
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.GetOrSetConnTimeOut())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		service.HandleError("mongodb", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		service.HandleError("mongodb", err)
	}

	return service.HandleSuccess("mongodb", nil)
}

//https://blog.logrocket.com/how-to-use-mongodb-with-go/
