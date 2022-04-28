package databases

import (
	"database/sql"
	"fmt"
	"strconv"
)

func MakeMysqlDbQuery(db *Database) map[string]string {
	uri, err := db.DSNMysql()
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	Mysqldb, err := sql.Open("mysql", uri)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	defer Mysqldb.Close()

	res, err := Mysqldb.Exec("SELECT 1")
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
	fmt.Println(status)
	return status
}

// https://golangbot.com/connect-create-db-mysql/
