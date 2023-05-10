package main

import (
	// "core"
	"core/config"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("JOSI :: START")
	// core.Start()
	log.Println(config.GetConfiguredServices())
	log.Println("JOSI :: FINISH")
}
