# go-watchdog
Go-watchdog is a web application observability tool built for Go, it exposes a status endpoint for application services like databases, caches, message-brokers, mails and even storages.


Go-watchdog is pretty much loosely-coupled hence allowing developers the ease to customize it to their use cases.


## Setting up Go-watchdog


``` go
	watchDogConfig := utils.WatchdogConfig{

		Databases: []databases.Database{

			{Type: "postgresql",
				Name:     "username",
				Addrs:    "db.elephantsql.com",
				Password: "password",
				Username: "username"},

			{Type: "postgresql",
				UriOnly: "",
			},

			{Type: "mongodb",
				UriOnly: "mongodb+srv://username:password@dbURL"},

			{Type: "couchbase",
				Addrs:    "",
				Username: "", Password: ""},
		},

		Caches: []caches.Cache{
			{Type: "redis", Addrs: "", Password: ""},
		},

		Storages: []storages.Storage{
			{Type: "aws", Region: os.Getenv("REGION"), BUCKET: os.Getenv("BUCKET")},
		},
	}

	utils.Register(&watchDogConfig)

```
