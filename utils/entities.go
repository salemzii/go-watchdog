package utils

import (
	"github.com/salemzii/go-watchdog/caches"
	"github.com/salemzii/go-watchdog/databases"
)

type WatchdogConfig struct {
	//Path      string               `json:"path"`
	Databases []databases.Database `json:"databases"`
	Caches    []caches.Cache       `json:"caches"`
}

func (wConfig *WatchdogConfig) HandleDbChecks() []databases.Database {
	return wConfig.Databases
}

func (wConfig *WatchdogConfig) HandleCacheChecks() []caches.Cache {
	return wConfig.Caches
}

func Register(conf *WatchdogConfig) {
	Config = *conf
}
