package databases

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/salemzii/go-watchdog/service"
)

func MakeSqlServerQuery(db *Database) service.ServiceCheck {
	uri, err := db.DSNSqlServer()
	if err != nil {
		return *service.HandleError("sql-server", err)
	}

	if db.Uri_Only() {
		uri = db.UriOnly
	}

	sqlDb, err := sql.Open("sqlserver", uri)
	if err != nil {
		return *service.HandleError("sql-server", err)
	}

	res, err := sqlDb.Exec("SELECT 1")
	if err != nil {
		return *service.HandleError("sql-server", err)
	}

	rows, err := res.RowsAffected()
	log.Println(rows)
	if err != nil {
		return *service.HandleError("sql-server", err)
	}
	return *service.HandleSuccess("sql-server", nil)

}
