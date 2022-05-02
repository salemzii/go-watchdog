package app

import "github.com/salemzii/go-watchdog/service"

func GetServiceCheck(res ...[]service.ServiceCheck) []service.ServiceCheck {
	serviceLs := []service.ServiceCheck{}

	for _, item := range res {

		for _, v := range item {
			serviceLs = append(serviceLs, v)
		}
	}
	/*
		var sc utils.ServiceCheck
			for _, item := range res {
				for _, v := range item {
					if _, ok := v["error"]; ok {
						sc.Error = v["err"]
						sc.Service = v["service"]
						sc.Status = v["status"]
					}
					sc.Service = v["service"]
					sc.Status = v["status"]

					serviceLs = append(serviceLs, sc)
				}
			}
	*/
	return serviceLs
}
