package app

import (
	"log"

	"github.com/salemzii/go-watchdog/service"
	"github.com/salemzii/go-watchdog/utils"
)

func AllCacheChecks() []service.ServiceCheck {
	checks, err := utils.GetCacheChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}
