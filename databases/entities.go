package databases

import (
	"errors"
	"log"
	"strings"
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

func (db *Database) GetDbDriver() {

	switch strings.ToLower(db.Type) {

	case "mysql":
		// call db driver
	case "postgresql":
		// call db driver
	case "sqlite":
		// call db driver
	case "oracle":
		// call db driver
	case "mongodb":
		// call db driver
	case "couchbase":
		// call db driver
	case "dynamodb":
		// call db driver
	default:
		log.Println("Db " + db.Type + " not supported")
	}
}
