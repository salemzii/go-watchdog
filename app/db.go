package app

import (
	"log"

	"github.com/salemzii/go-watchdog/utils"
)

func AllDbChecks() []map[string]string {
	checks, err := utils.GetDatabaseChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}
