package main

import (
	"github.com/salemzii/go-watchdog/app"
	"github.com/salemzii/go-watchdog/messagequeues"
	"github.com/salemzii/go-watchdog/utils"
)

func main() {

	watchDogConfig := utils.WatchdogConfig{

		/*
			Databases: []databases.Database{

				{Type: "sqlite3", Name: "test2.db"},
				{Type: "sqlite3", Name: "test.db"},
				{Type: "mongodb", Name: "taskdb", Addrs: "127.0.0.1:27017"},
				{Type: "sql-server", Username: "SA", Addrs: "127.0.0.1:1433",
					Password: os.Getenv("SQLSERVERPSWD"), Name: "TestDb"},
				{Type: "postgresql", Name: "postgres", Addrs: "localhost", Username: "postgres",
					Password: os.Getenv("PG_LOCAL_PSWD")},

				{Type: "postgresql",
					UriOnly: os.Getenv("PG_URI"),
				},

				{Type: "mongodb",
					UriOnly: os.Getenv("MONGO_URI")},

				{Type: "couchbase",
					Addrs:    os.Getenv("COUCHBASE_URL"),
					Username: "salemododa2@gmail.com", Password: os.Getenv("COUCHBASE_PSWD")},
			},

			Caches: []caches.Cache{
				{Type: "memcached", Addrs: "localhost:11211"},
				{Type: "redis", Addrs: "127.0.0.1:6379"},
				{Type: "redis", Addrs: os.Getenv("REDIS_URI"), Password: os.Getenv("REDIS_PSWD")},
			},

			Storages: []storages.Storage{
				{Type: "aws", Region: os.Getenv("REGION"), BUCKET: os.Getenv("BUCKET")},
				{Type: "aws", Region: os.Getenv("REGION"), BUCKET: os.Getenv("BUCKET")},
			},

		*/
		MsgQueues: []messagequeues.MsgQueue{
			{Type: "awsmq", UriOnly: "amqps://b-4de5a854-a3c3-43e3-9f3d-d8fe82060e84-1.mq.us-east-1.amazonaws.com:5671"},
			{Type: "rabbitmq", Addrs: "localhost:5672", Username: "guest", Password: "guest"},
			{Type: "rabbitmq", UriOnly: ""},
		},
	}

	utils.Register(&watchDogConfig)

	app.RunGinserver()

}
