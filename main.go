package main

import (
	"NewApp/config"
	"NewApp/database"
	"NewApp/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	database.Init()
	server.Init()
}
