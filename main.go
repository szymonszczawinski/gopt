package main

import (
	"embed"
	"gosi/core"
	"gosi/core/config"
	"gosi/core/http"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//go:embed static/*
var publicDir embed.FS

func init() {
	godotenv.Load()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("GOSI :: START")
	staticContent := http.StaticContent{
		PublicDir: publicDir,
	}
	cla := parseCLA(os.Args)
	core.Start(cla, staticContent)

	log.Println("GOSI :: FINISH")
	// sqlite.GetSqliteRepository()
	// RunTests()
}

func parseCLA(args []string) map[string]any {
	cla := map[string]any{}
	cla[config.RUN_MODE] = config.RUN_MODE_DEV
	if len(args) > 1 {
		args = args[1:]
		if len(args)%2 == 1 {
			log.Println("Incorrect  number of command line parameters", args)
			return cla
		}
		for i := 0; i < len(args); i += 2 {
			if args[i] == config.RUN_MODE {
				cla[config.RUN_MODE] = args[i+1]
			}
		}

	}

	return cla
}
