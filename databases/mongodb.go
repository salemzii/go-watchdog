package databases

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeMongodbQueryCheck(db *Database) map[string]string {

	uri, err := db.DSNMongoDb()
	if err != nil {
		status := handleDberr(err)
		return status
	}
	if db.Uri_Only() {
		uri = db.UriOnly
	}

	ctx, cancel := context.WithTimeout(context.Background(), db.GetOrSetConnTimeOut())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		status := handleDberr(err)
		return status
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		status := handleDberr(err)
		return status
	}

	status := map[string]string{
		"status": "ok",
	}
	fmt.Println(status)
	return status
}

//https://blog.logrocket.com/how-to-use-mongodb-with-go/
