# go-watchdog
Go-watchdog is a web application observability tool built for Go, it exposes a status endpoint for application services like databases, caches, message-brokers, mails and storages.


## Setting up Go-watchdog

Go-watchdog is pretty much loosely-coupled hence allowing developers the ease to customize it to their use cases.
The bases schema for configuring watchdog, is the `WatchdogConfig struct`,
	
	# The watchdog Struct
	
	import(
		"github.com/salemzii/go-watchdog/caches"
		"github.com/salemzii/go-watchdog/databases"
		storages "github.com/salemzii/go-watchdog/storages"
	)
	
	type WatchdogConfig struct {
		//Path      string               `json:"path"`
		Databases []databases.Database `json:"databases"`
		Caches    []caches.Cache       `json:"caches"`
		Storages  []storages.Storage   `json:"storages"`
	}
	
	
	

``` go
	# Initializing WatchDog
	
	import (
		app "github.com/salemzii/go-watchdog/app"
		"github.com/salemzii/go-watchdog/caches"
		"github.com/salemzii/go-watchdog/databases"
		storages "github.com/salemzii/go-watchdog/storages"
		"github.com/salemzii/go-watchdog/utils"
	)

	watchDogConfig := utils.WatchdogConfig{

		Databases: []databases.Database{

			{Type: "postgresql",
				Name:     "username",
				Addrs:    "db.elephantsql.com",
				Password: "password",
				Username: "username"},

			{Type: "mongodb",
				UriOnly: "mongodb+srv://username:password@dbURL"},
		},

		Caches: []caches.Cache{
			{Type: "redis", Addrs: "", Password: ""},
		},

		Storages: []storages.Storage{
			{Type: "aws", Region: "", BUCKET: ""},
		},
	}


```
Since, each service (i.e databases, storages etc) requires distinct parameters to make connection to it's server, the watchdog struct fields generally represents a slice of particular service type.


## Registering Watchdog
Once you've configured the services you want to watch as specified above, you can call the `utils.Register()` function to make your watchdog config available to the the watchdog "ServiceCheckers".
Usually you can register the config by passing it's memory address to the register function:
	```
		func Register(conf *WatchdogConfig) {
			Config = *conf
		}```
		
	"utils.Register(&watchDogConfig)"
	
	

## Databases

The watchdog database struct lets you config your db's credentials;

	Type: The type of database you're configuring.
	Name: The name of the database you're connecting to.
	Addrs: The uri to the database's server.
	Username: The db's username.
	Password: The db's password.
	Timeout: number of seconds to wait before timeout error, Watchdog sets the default to 10secs.
	UriOnly: In Cases where db's credentials are already preconfigured to a single address/URL. If this value is set, watchdog automatically makes 		connection to the db's server using the credentials specified in the uri.

	``` go 
		type Database struct {
			Type     string `json:"type"`
			Name     string `json:"name"`
			Username string `json:"username"`
			Password string `json:"password"`
			Addrs    string `json:"addrs"`
			Timeout  uint   `json:"timeout"`
			UriOnly  string `json:"uri_only"` // when db credentials are already preconfigured to a single address/URL
			Uri      string `json:"uri"`     
		}```
		
Supported Databases:
Currently go-watchdog supports just a handful of databases services, we are working endlessly to add support for more databases, you can also do well to contribute as any effort is appreciated.

	```
		"SQL":   {"mysql", "sqlite", "postgresql", "oracle", "sql-server"},
		"NO-SQL": {"mongodb", "couchbase", "dynamodb", "cockroachDB"},
	```
