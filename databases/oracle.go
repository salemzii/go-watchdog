package databases

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/salemzii/go-watchdog/service"
)

func MakeOracleDbQuery(db *Database) service.ServiceCheck {
	uri, err := db.DSNOracle()
	if err != nil {
		return service.HandleError("oracledb", err)
	}
	Oracledb, err := sql.Open("godror", uri)
	if err != nil {
		return service.HandleError("oracledb", err)
	}
	defer Oracledb.Close()

	res, err := Oracledb.Exec("SELECT 1")
	if err != nil {
		return service.HandleError("oracledb", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return service.HandleError("oracledb", err)
	}
	status := map[string]string{
		"service":       "oracle",
		"status":        "ok",
		"rows_affected": strconv.Itoa(int(rows)),
	}
	fmt.Println(status)
	return service.HandleSuccess("oracledb", nil)
}

//https://blogs.oracle.com/developers/post/how-to-connect-a-go-program-to-oracle-database-using-godror
//https://medium.com/venturenxt/install-oracle-database-12c-on-ubuntu-16-04-c081d51c0f9d
