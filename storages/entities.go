package storages

import (
	"log"
	"strings"
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

func (st *Storage) GetStorageDriver() map[string]string {

	switch strings.ToLower(st.Type) {
	case "aws":
		return AwsStorageCheck(st)
	case "linode":
		return map[string]string{}
	case "digitalocean":
		return map[string]string{}
	default:
		log.Fatal("Storahe " + st.Type + " not supported")
	}

	return map[string]string{}
}
