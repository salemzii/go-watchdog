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
	"github.com/salemzii/go-watchdog/service"
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

func (db *Database) GetDbDriver() service.ServiceCheck {

	switch strings.ToLower(db.Type) {

	case "mysql":
		return MakeMysqlDbQuery(db)
	case "postgresql":
		return MakePostgresDbQuery(db)
	case "sqlite3":
		return MakeSqliteQueryCheck(db)
	case "oracle":
		return MakeOracleDbQuery(db)
	case "mongodb":
		return MakeMongodbQueryCheck(db)
	case "couchbase":
		return MakeCouchDbQueryCheck(db)
	case "dynamodb":
		// call db driver
	default:
		log.Fatal("Db " + db.Type + " not supported")

	}
	return service.ServiceCheck{}
}

func (db *Database) GetOrSetConnTimeOut() time.Duration {
	if db.Timeout != 0 {
		return time.Duration(db.Timeout * uint(time.Second))
	}
	db.Timeout = uint(10)
	return time.Duration(db.Timeout * uint(time.Second))
}

func handleDberr(service string, err error) map[string]string {
	status := map[string]string{
		"status":  "Failed",
		"error":   err.Error(),
		"service": service,
	}
	return status
}

// GENERICS
// Database
// connection-pooling
