package databases

import (
	"database/sql"
	"fmt"
	"strconv"
)

func MakePostgresDbQuery(db *Database) map[string]string {
	uri, err := db.GetConnString()
	if err != nil {
		status := handleDberr("postgresql", err)
		fmt.Println(status)
		return status
	}

	if db.Uri_Only() {
		uri = db.UriOnly
	}

	postgresDb, err := sql.Open("postgres", uri)
	if err != nil {
		status := handleDberr("postgresql", err)
		fmt.Println(status)
		return status
	}

	res, err := postgresDb.Exec("SELECT 1")
	if err != nil {
		status := handleDberr("postgresql", err)
		fmt.Println(status)
		return status
	}

	rows, err := res.RowsAffected()
	if err != nil {
		status := handleDberr("postgresql", err)
		fmt.Println(status)
		return status
	}
	status := map[string]string{
		"service":       "postgresql",
		"status":        "ok",
		"rows_affected": strconv.Itoa(int(rows)),
	}
	return status
}

// https://blog.logrocket.com/building-simple-app-go-postgresql/
// https://stackoverflow.com/questions/10845998/i-forgot-the-password-i-entered-during-postgres-installation
