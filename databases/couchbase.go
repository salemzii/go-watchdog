package databases

import (
	"fmt"

	"github.com/couchbase/gocb"
)

func MakeCouchDbQueryCheck(db *Database) map[string]string {
	uri, err := db.DSNCouchbase()

	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	cluster, err := gocb.Connect(uri)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	bucket, err := cluster.OpenBucket("default", db.Password)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	r, err := bucket.Ping([]gocb.ServiceType{gocb.N1qlService})
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	fmt.Println(r)
	status := map[string]string{
		"status": "ok",
		"report": "",
	}
	fmt.Println(status)
	return status
}
