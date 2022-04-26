package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/orders/", MyOrders)

	log.Println("Starting server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
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
