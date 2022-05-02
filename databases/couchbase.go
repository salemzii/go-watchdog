package databases

import (
	"log"

	"github.com/couchbase/gocb/v2"
	"github.com/salemzii/go-watchdog/service"
)

func MakeCouchDbQueryCheck(db *Database) service.ServiceCheck {
	uri, err := db.DSNCouchbase()

	if err != nil {
		return service.HandleError("couchbase", err)
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
		return service.HandleError("couchbase", err)
	}
	bucket := cluster.Bucket("default")

	r := bucket.Name()

	log.Println(r)
	return service.HandleSuccess("couchbase", nil)
}
