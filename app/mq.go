package app

import (
	"log"

	"github.com/salemzii/go-watchdog/service"
	"github.com/salemzii/go-watchdog/utils"
)

func AllMqChecks() *[]service.ServiceCheck {
	checks, err := utils.GetMQChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}