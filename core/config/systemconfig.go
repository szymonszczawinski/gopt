package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	ROOT_DIR      string = ".."
	CONFIG_DIR    string = ROOT_DIR + "/config"
	RUN_MODE      string = "-runmode"
	RUN_MODE_DEV  string = "dev"
	RUN_MODE_PLUG string = "plug"
)

var configuredServices ServiceConfigItems

func GetConfiguredServices() (*ServiceConfigItems, error) {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println(cwd)
	// Open our jsonFile
	jsonFile, err := os.Open(CONFIG_DIR + "/services_config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Successfully Opened services_config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &configuredServices)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return &configuredServices, nil
}
