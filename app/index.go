package app

import "github.com/salemzii/go-watchdog/service"

func GetServiceCheck(res ...[]service.ServiceCheck) []service.ServiceCheck {
	serviceLs := []service.ServiceCheck{}

	for _, item := range res {

		for _, v := range item {
			serviceLs = append(serviceLs, v)
		}
	}
	return serviceLs
}
