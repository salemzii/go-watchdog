package utils

import (
	"github.com/salemzii/go-watchdog/caches"
	"github.com/salemzii/go-watchdog/databases"
	"github.com/salemzii/go-watchdog/storages"
)

type WatchdogConfig struct {
	//Path      string               `json:"path"`
	Databases []databases.Database `json:"databases"`
	Caches    []caches.Cache       `json:"caches"`
	Storages  []storages.Storage   `json:"storages"`
}

func (wConfig *WatchdogConfig) HandleDbChecks() []databases.Database {
	return wConfig.Databases
}

func (wConfig *WatchdogConfig) HandleCacheChecks() []caches.Cache {
	return wConfig.Caches
}
func (wConfig *WatchdogConfig) HandleStorageCheck() []storages.Storage {
	return wConfig.Storages
}

func Register(conf *WatchdogConfig) {
	Config = *conf
}
