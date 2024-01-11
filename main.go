package main

import (
	"embed"
	"fmt"
	"gosi/core"
	"gosi/core/config"
	domain "gosi/core/domain/issue"
	"gosi/core/http"
	"log"
	"os"
	"reflect"
	"runtime"

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
			} else if args[i] == config.INIT_DB {
				cla[config.INIT_DB] = args[i+1]
			}
		}

	}

	return cla
}

func testStructSize() {
	typ := reflect.TypeOf(domain.Issue{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size())
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}

	// allStats := []stats{}
	// for i := 0; i < 100000000; i++ {
	// 	allStats = append(allStats, stats{})
	// }

	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
