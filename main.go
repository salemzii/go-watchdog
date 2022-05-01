package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/salemzii/go-watchdog/app"
	"github.com/salemzii/go-watchdog/storages"
	"github.com/salemzii/go-watchdog/utils"
)

func main() {

	watchDogConfig := utils.WatchdogConfig{

		/*
			Databases: []databases.Database{

				{Type: "sqlite3", Name: "test2.db"},
				{Type: "sqlite3", Name: "test.db"},
				{Type: "mongodb", Name: "taskdb", Addrs: "127.0.0.1:27017"},
				{Type: "postgresql", Name: "postgres", Addrs: "localhost", Username: "postgres", Password: "auth1234"},

				{Type: "postgresql",
					Name:     "tfgrwusb",
					Addrs:    "queenie.db.elephantsql.com",
					Password: "MwZ8sT4H0_8575ybn2yaTz3h3ImAlp40",
					Username: "tfgrwusb"},

				{Type: "postgresql",
					UriOnly: "postgres://tfgrwusb:MwZ8sT4H0_8575ybn2yaTz3h3ImAlp40@queenie.db.elephantsql.com/tfgrwusb",
				},

				{Type: "mongodb",
					UriOnly: "mongodb+srv://salem:auth1234@cluster0.8qw1s.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"},

				{Type: "couchbase",
					Addrs:    "cb.lus1jnhsaeag2nl4.cloud.couchbase.com",
					Username: "salemododa2@gmail.com", Password: "Auth1234#"},
			},

			Caches: []caches.Cache{
				{Type: "memcached", Addrs: "localhost:11211"},
				{Type: "redis", Addrs: "127.0.0.1:6379"},
				{Type: "redis", Addrs: "redis-15719.c242.eu-west-1-2.ec2.cloud.redislabs.com:15719", Password: "38rKjb8yOD7YI2OodiAoFdrMZQTIBIYl"},
			},
		*/
		Storages: []storages.Storage{
			{Type: "aws", Region: os.Getenv("REGION"), BUCKET: os.Getenv("BUCKET")},
		},
	}

	utils.Register(&watchDogConfig)
	fmt.Println(app.AllDbChecks())
	fmt.Println(app.AllCacheChecks())
	fmt.Print(app.AllStorageChecks())

	/*
		http.HandleFunc("/orders/", MyOrders)
		fmt.Println(app.AllDbChecks())
		log.Println("Starting server on port :8000")
		log.Fatal(http.ListenAndServe(":8000", nil))

	*/
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

func MyOrders(wr http.ResponseWriter, req *http.Request) {
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

	t.Execute(wr, par)
}

// https://blog.logrocket.com/using-golang-templates/
