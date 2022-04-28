package databases

import (
	"fmt"

	"github.com/couchbase/gocb/v2"
)

func MakeCouchDbQueryCheck(db *Database) map[string]string {
	uri, err := db.DSNCouchbase()

	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	cluster, err := gocb.Connect(uri, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: db.Username,
			Password: db.Password,
		},
		SecurityConfig: gocb.SecurityConfig{
			TLSSkipVerify: true,
		},
	})
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	bucket := cluster.Bucket("default")

	r := bucket.Name()

	fmt.Println(r)
	status := map[string]string{
		"status": "ok",
		"report": r,
	}
	fmt.Println(status)
	return status
}
