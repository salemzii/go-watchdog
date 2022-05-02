package service

type ServiceCheck struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

func HandleError(service string, err error) ServiceCheck {
	return ServiceCheck{Service: service, Error: err.Error(), Status: "Failed"}
}
func HandleSuccess(service string, err error) ServiceCheck {
	return ServiceCheck{Service: service, Error: err.Error(), Status: "Success"}
}
