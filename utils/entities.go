package utils

import (
	"github.com/salemzii/go-watchdog/databases"
)

type WatchdogConfig struct {
	//Path      string               `json:"path"`
	Databases []databases.Database `json:"databases"`
	// Caches []
}

func (wConfig *WatchdogConfig) HandleDbChecks() []databases.Database {
	return wConfig.Databases
}

func Register(conf *WatchdogConfig) {
	Config = *conf
}
