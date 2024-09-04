package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"log/slog"
	"os"
)

const (
	CONFIG_DIR       string = "/config"
	HTTP_SERVER_PORT string = "httpServerPort"
)

func GetClientConfig() (map[string]any, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slog.Info(cwd)
	// Open our jsonFile
	jsonFile, err := os.Open(cwd + "/" + CONFIG_DIR + "/client_config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		slog.Info("error opening client config", "err", err)
		return nil, err
	}
	slog.Info("Successfully Opened services_config.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	configMap := map[string]any{}
	json.Unmarshal(byteValue, &configMap)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return configMap, nil
}
