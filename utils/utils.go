package utils

import "github.com/salemzii/go-watchdog/service"

var Config WatchdogConfig

// go build -v *.go && ./main
// sudo systemctl start mongod

func GetDatabaseChecks() (checks []service.ServiceCheck, err error) {
	arg := Config.Databases
	allDbChecks := []service.ServiceCheck{}
	for i := 0; i < len(arg); i++ {
		status := arg[i].GetDbDriver()
		allDbChecks = append(allDbChecks, status)
	}
	return allDbChecks, nil
}

func GetCacheChecks() (checks []service.ServiceCheck, err error) {
	arg := Config.Caches
	allCacheChecks := []service.ServiceCheck{}

	for i := 0; i < len(arg); i++ {
		status := arg[i].GetCacheDriver()
		allCacheChecks = append(allCacheChecks, status)
	}
	return allCacheChecks, nil
}

func GetStorageChecks() (checks []service.ServiceCheck, err error) {
	arg := Config.Storages
	allCacheChecks := []service.ServiceCheck{}

	for i := 0; i < len(arg); i++ {
		status := arg[i].GetStorageDriver()
		allCacheChecks = append(allCacheChecks, status)
	}
	return allCacheChecks, nil
}
