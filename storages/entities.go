package storages

import (
	"log"
	"strings"

	"github.com/salemzii/go-watchdog/service"
)

type Storage struct {
	Type    string `json:"type"`
	Region  string `json:"region"`
	BUCKET  string `json:"bucket"`
	UriOnly string `json:"urionly"`
}

func (st *Storage) getAwsDns()          {}
func (st *Storage) getLinodeDns()       {}
func (st *Storage) getDigitalOceanDns() {}

func (st *Storage) GetStorageDriver() service.ServiceCheck {

	switch strings.ToLower(st.Type) {
	case "aws":
		return AwsStorageCheck(st)
	case "linode":
		return service.ServiceCheck{}
	case "digitalocean":
		return service.ServiceCheck{}
	default:
		log.Fatal("Storahe " + st.Type + " not supported")
	}

	return service.ServiceCheck{}
}

func handleStorageErr(service string, err error) map[string]string {

	return map[string]string{"status": "Fail", "service": service, "error": err.Error()}
}
