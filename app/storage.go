package app

import (
	"log"

	"github.com/salemzii/go-watchdog/utils"
)

func AllStorageChecks() []map[string]string {
	checks, err := utils.GetStorageChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}
