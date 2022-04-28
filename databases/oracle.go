package databases

import (
	"database/sql"
	"fmt"
	"strconv"
)

func MakeOracleDbQuery(db *Database) map[string]string {
	uri, err := db.DSNOracle()
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	Oracledb, err := sql.Open("godror", uri)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	defer Oracledb.Close()

	res, err := Oracledb.Exec("SELECT 1")
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

//https://blogs.oracle.com/developers/post/how-to-connect-a-go-program-to-oracle-database-using-godror
//https://medium.com/venturenxt/install-oracle-database-12c-on-ubuntu-16-04-c081d51c0f9d
