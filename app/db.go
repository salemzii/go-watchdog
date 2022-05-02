package app

import (
	"log"

	"github.com/salemzii/go-watchdog/service"
	"github.com/salemzii/go-watchdog/utils"
)

func AllDbChecks() []service.ServiceCheck {
	checks, err := utils.GetDatabaseChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}
