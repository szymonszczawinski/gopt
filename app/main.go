package main

import (
	"core"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("JOSI :: START")
	core.Start()
	log.Println("JOSI :: FINISH")
}
