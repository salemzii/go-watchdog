package app

import (
	"log"

	"github.com/salemzii/go-watchdog/utils"
)

func AllCacheChecks() []map[string]string {
	checks, err := utils.GetCacheChecks()
	if err != nil {
		log.Println(err)
	}
	return checks
}
