package databases

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/salemzii/go-watchdog/service"
)

func MakeSqliteQueryCheck(db *Database) service.ServiceCheck {

	sqldb, err := sql.Open("sqlite3", db.Name)
	if err != nil {
		return *service.HandleError("sqlite3", err)
	}

	defer sqldb.Close()
	res, err := sqldb.Exec("SELECT 1")

	if err != nil {

		return *service.HandleError("sqlite3", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {

		return *service.HandleError("sqlite3", err)
	}

	status := map[string]string{
		"service":       "sqlite3",
		"status":        "ok",
		"rows_affected": strconv.Itoa(int(rows)),
	}
	fmt.Println(status)
	return *service.HandleSuccess("sqlite", nil)
}
