package utils

var Config WatchdogConfig

// go build -v *.go && ./main
// sudo systemctl start mongod

/*
func init() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	workingDir := filepath.Dir(folderPath)
	os.Chdir(workingDir)
	config, err := os.Open("watchdogConfig.json")
	if err != nil {
		log.Println("Starting AutoReloader with config file .......")
	}

	defer config.Close()
	configByte, err := ioutil.ReadAll(config)
	if err != nil {
		log.Println("Error loading AutoReloaderConfig.json")
	}
	json.Unmarshal(configByte, &Config)
	fmt.Println(Config)
}
*/

func GetDatabaseChecks() (checks []map[string]string, err error) {
	arg := Config.Databases
	allDbChecks := []map[string]string{}
	for i := 0; i < len(arg); i++ {
		status := arg[i].GetDbDriver()
		allDbChecks = append(allDbChecks, status)
	}
	return allDbChecks, nil
}
