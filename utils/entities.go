package utils

import "github.com/salemzii/go-watchdog/databases"

type WatchdogConfig struct {
	Databases []databases.Database `json:"databases"`
	// Caches []
}
