package main

import (
	"html/template"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	app "github.com/salemzii/go-watchdog/app"
	"github.com/salemzii/go-watchdog/caches"
	"github.com/salemzii/go-watchdog/databases"
	storages "github.com/salemzii/go-watchdog/storages"
	"github.com/salemzii/go-watchdog/utils"
)

func main() {

	router := gin.Default()

	watchDogConfig := utils.WatchdogConfig{

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
	}

	utils.Register(&watchDogConfig)

	router.GET("/storages", app.GinLookUp)
	router.GET("/order", MyOrders)

	router.Run()
}

type Order struct {
	Number  uint
	Details string
	Status  string
	Date    time.Time
	Total   float64
}
type OrderLs struct {
	Orders []Order
}

func MyOrders(c *gin.Context) {
	order1 := &Order{
		Number:  1001,
		Details: "Kring New Fit kitchen chair, couch + PU, brown",
		Status:  "Delivered",
		Date: time.Date(2022, time.April,
			11, 21, 34, 01, 0, time.UTC),
		Total: 125.00,
	}

	order2 := &Order{
		Number:  1004,
		Details: "Sky blue, T-shirt, lg",
		Status:  "Processing",
		Date: time.Date(2022, time.April,
			11, 27, 34, 01, 0, time.UTC),
		Total: 22.00,
	}

	par := OrderLs{Orders: []Order{*order1, *order2}}
	t, _ := template.ParseFiles("template/index.html")

	t.Execute(c.Writer, par)
}

// https://blog.logrocket.com/using-golang-templates/
