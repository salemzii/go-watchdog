package databases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/couchbase/gocb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var supportedDbs = map[string][]string{
	"sql":   {"mysql", "sqlite", "postgresql", "oracle"},
	"nosql": {"mongodb", "couchbase", "dynamodb"},
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
	return fmt.Sprintf("%s://%s", db.Type, db.Addrs), nil
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

// GENERICS
// Database

func handleDberr(err error) map[string]string {
	status := map[string]string{
		"status": "Fail",
		"error":  err.Error(),
	}
	return status
}

func MakeSqliteQueryCheck(db *Database) map[string]string {

	sqldb, err := sql.Open("sqlite3", db.Name)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	defer sqldb.Close()
	res, err := sqldb.Exec("SELECT 1")

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

func MakeMongodbQueryCheck(db *Database) map[string]string {

	/*
		mongoDialInfo := &mgo.DialInfo{Addrs: []string{db.Addrs},
			Timeout:  db.GetOrSetConnTimeOut(),
			Database: db.Name,
			Username: db.Username,
			Password: db.Password,
		}

		mongoSession, err := mgo.DialWithInfo(mongoDialInfo)
		if err != nil {
			status := handleDberr(err)
			fmt.Println(status)
			return status
		}

	*/
	uri, err := db.DSNMongoDb()
	if err != nil {
		status := handleDberr(err)
		return status
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		status := handleDberr(err)
		return status
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		status := handleDberr(err)
		return status
	}

	status := map[string]string{
		"status": "ok",
	}
	fmt.Println(status)
	return status
}

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

func MakeCouchDbQueryCheck(db *Database) map[string]string {
	uri, err := db.DSNCouchbase()

	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	cluster, err := gocb.Connect(uri)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}
	bucket, err := cluster.OpenBucket("default", db.Password)
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	r, err := bucket.Ping([]gocb.ServiceType{gocb.N1qlService})
	if err != nil {
		status := handleDberr(err)
		fmt.Println(status)
		return status
	}

	fmt.Println(r)
	status := map[string]string{
		"status": "ok",
		"report": "",
	}
	fmt.Println(status)
	return status
}
