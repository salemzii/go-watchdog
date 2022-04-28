package databases

import (
	"database/sql"
	"fmt"
	"strconv"
)

func MakePostgresDbQuery(db *Database) map[string]string {
	uri, err := db.GetConnString()
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	postgresDb, err := sql.Open("postgres", uri)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	res, err := postgresDb.Exec("SELECT 1")
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	rows, err := res.RowsAffected()
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	status := map[string]string{
		"status":        "ok",
		"rows_affected": strconv.Itoa(int(rows)),
	}
	return status
}

// https://blog.logrocket.com/building-simple-app-go-postgresql/
