package config

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
)

const (
	ROOT_DIR      string = ".."
	CONFIG_DIR    string = ROOT_DIR + "/config"
	RUN_MODE      string = "-runmode"
	RUN_MODE_DEV  string = "dev"
	RUN_MODE_PLUG string = "plug"

	// application cmd arg to drop and recreate database schema
	INIT_DB       string = "-initdb"
	INIT_DB_TRUE  string = "true"
	INIT_DB_FALSE string = "false"
)

var (
	configuredServices ServiceConfigItems

	systemStartParameters map[string]any
)

func InitSystemConfiguration(config map[string]any) {
	systemStartParameters = config
}

func GetSystemConfig(key string) any {
	return systemStartParameters[key]
}

func GetConfiguredServices() (*ServiceConfigItems, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slog.Info(cwd)
	// Open our jsonFile
	jsonFile, err := os.Open(CONFIG_DIR + "/services_config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		slog.Info("get configures services", "err", err)
		return nil, err
	}
	slog.Info("Successfully Opened services_config.json")
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &configuredServices)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return &configuredServices, nil
}
