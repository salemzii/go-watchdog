package databases

import (
	"database/sql"
	"log"

	"github.com/salemzii/go-watchdog/service"
)

func MakePostgresDbQuery(db *Database) service.ServiceCheck {
	uri, err := db.GetConnString()
	if err != nil {
		return service.HandleError("postgresql", err)
	}

	if db.Uri_Only() {
		uri = db.UriOnly
	}

	postgresDb, err := sql.Open("postgres", uri)
	if err != nil {
		return service.HandleError("postgresql", err)
	}

	res, err := postgresDb.Exec("SELECT 1")
	if err != nil {
		return service.HandleError("postgresql", err)
	}

	rows, err := res.RowsAffected()
	log.Println(rows)
	if err != nil {
		return service.HandleError("postgresql", err)
	}
	return service.HandleSuccess("postgresql", nil)
}

// https://blog.logrocket.com/building-simple-app-go-postgresql/
// https://stackoverflow.com/questions/10845998/i-forgot-the-password-i-entered-during-postgres-installation
