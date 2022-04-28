package databases

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var supportedDbs = map[string][]string{
	"sql":   {"mysql", "sqlite", "postgresql", "oracle"},
	"nosql": {"mongodb", "couchbase", "dynamodb", "cockroachDB"},
}

type Databases struct {
	DatabaseClusters []Database `json:"databases"`
}

type Database struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Addrs    string `json:"addrs"`
	Timeout  uint   `json:"timeout"`
	UriOnly  string `json:"uri_only"` // when db credentials are already preconfigured to a single address/URL
	Uri      string `json:"uri"`      // similar to UriOnly but without the password/auth info to the authenticate the uri
}

func (db *Database) Uri_Only() bool {
	return db.UriOnly != ""
}

func (db *Database) GetConnString() (str string, err error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", db.Type, db.Username, db.Password, db.Addrs, db.Name)

	return connStr, nil
}

func (db *Database) DSNMysql() (str string, err error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", db.Username, db.Password, db.Addrs, db.Name)
	return connStr, nil
}

func (db *Database) DSNOracle() (str string, err error) {
	connStr := fmt.Sprintf("%s/%s@%s", db.Username, db.Password, db.Addrs)
	return connStr, nil
}
func (db *Database) DSNCouchbase() (str string, err error) {
	return fmt.Sprintf("couchbase://%s", db.Addrs), nil
}

func (db *Database) DSNMongoDb() (str string, err error) {
	if db.Username != "" && db.Password != "" {
		return fmt.Sprintf("%s+srv://%s:%s@%s/%s?retryWrites=true&w=majority", db.Type, db.Username, db.Password, db.Addrs, db.Name), nil
	}
	return fmt.Sprintf("%s://%s/%s", db.Type, db.Addrs, db.Name), nil
}

func (db *Database) GetDbSupported() (supported bool, err error) {
	for _, val := range supportedDbs {
		for _, value := range val {
			if db.Type == value {
				return true, nil
			}
		}
	}
	return false, errors.New("Db " + db.Type + " not supported")
}

func (db *Database) GetDbDriver() map[string]string {

	switch strings.ToLower(db.Type) {

	case "mysql":
		status := MakeMysqlDbQuery(db)
		return status
	case "postgresql":
		status := MakePostgresDbQuery(db)
		return status
	case "sqlite3":
		status := MakeSqliteQueryCheck(db)
		return status
	case "oracle":
		status := MakeOracleDbQuery(db)
		return status
	case "mongodb":
		status := MakeMongodbQueryCheck(db)
		return status
	case "couchbase":
		status := MakeCouchDbQueryCheck(db)
		return status
	case "dynamodb":
		// call db driver
	default:
		log.Fatal("Db " + db.Type + " not supported")

	}
	return map[string]string{}
}

func (db *Database) GetOrSetConnTimeOut() time.Duration {
	if db.Timeout != 0 {
		return time.Duration(db.Timeout * uint(time.Second))
	}
	db.Timeout = uint(10)
	return time.Duration(db.Timeout * uint(time.Second))
}

func handleDberr(err error) map[string]string {
	status := map[string]string{
		"status": "Fail",
		"error":  err.Error(),
	}
	return status
}

// GENERICS
// Database
// connection-pooling
