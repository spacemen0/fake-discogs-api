package main

import (
	"fake-discogs-api/config"
	"fake-discogs-api/database"
	"fake-discogs-api/server"
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
