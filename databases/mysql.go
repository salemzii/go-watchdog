package databases

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/salemzii/go-watchdog/service"
)

func MakeMysqlDbQuery(db *Database) service.ServiceCheck {
	uri, err := db.DSNMysql()
	if err != nil {
		return *service.HandleError("mysql", err)
	}

	if db.Uri_Only() {
		uri = db.UriOnly
	}

	Mysqldb, err := sql.Open("mysql", uri)
	if err != nil {
		return *service.HandleError("mysql", err)
	}
	defer Mysqldb.Close()

	res, err := Mysqldb.Exec("SELECT 1")
	if err != nil {
		return *service.HandleError("mysql", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return *service.HandleError("mysql", err)
	}
	status := map[string]string{
		"service":       "mysql",
		"status":        "ok",
		"rows_affected": strconv.Itoa(int(rows)),
	}
	log.Println(status)
	return *service.HandleSuccess("mysql", nil)
}

// https://golangbot.com/connect-create-db-mysql/
