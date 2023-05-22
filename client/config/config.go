package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const CONFIG_DIR string = "/config"
const HTTP_SERVER_PORT string = "httpServerPort"

func GetClientConfig() (map[string]any, error) {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println(cwd)
	// Open our jsonFile
	jsonFile, err := os.Open(cwd + "/" + CONFIG_DIR + "/client_config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Successfully Opened services_config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	configMap := map[string]any{}
	json.Unmarshal(byteValue, &configMap)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return configMap, nil
}
