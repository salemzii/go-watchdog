package utils

import (
	"github.com/salemzii/go-watchdog/databases"
)

type WatchdogConfig struct {
	//Path      string               `json:"path"`
	Databases []databases.Database `json:"databases"`
	// Caches []
}

func HandleDb() {

}
