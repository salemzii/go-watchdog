package databases

import (
	"database/sql"
	"fmt"
	"strconv"
)

func MakeSqliteQueryCheck(db *Database) map[string]string {

	sqldb, err := sql.Open("sqlite3", db.Name)
	if err != nil {
		status := handleDberr("sqlite3", err)
		fmt.Println(status)
		return status
	}

	defer sqldb.Close()
	res, err := sqldb.Exec("SELECT 1")

	if err != nil {
		status := handleDberr("sqlite3", err)
		fmt.Println(status)
		return status
	}

	rows, err := res.RowsAffected()
	if err != nil {
		status := handleDberr("sqlite3", err)
		fmt.Println(status)
		return status
	}

	status := map[string]string{
		"service":       "sqlite3",
		"status":        "ok",
		"rows_affected": strconv.Itoa(int(rows)),
	}
	fmt.Println(status)
	return status
}
