package app

import (
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/salemzii/go-watchdog/service"
)

func GinLookUp(c *gin.Context) {
	type ServiceCheckTemplate struct {
		ServiceChecks []service.ServiceCheck `json:"service_checks"`
	}
	serviceChecks := GetServiceCheck(*AllDbChecks(), *AllCacheChecks(), *AllStorageChecks())
	par := ServiceCheckTemplate{ServiceChecks: serviceChecks}
	log.Println(par)
	t, _ := template.ParseFiles("template/utils_templating.html")

	t.Execute(c.Writer, par)
}
