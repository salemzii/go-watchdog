package app

import (
	"log"

	"github.com/salemzii/go-watchdog/service"
	"github.com/salemzii/go-watchdog/utils"
)

func AllStorageChecks() []service.ServiceCheck {
	checks, err := utils.GetStorageChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}
