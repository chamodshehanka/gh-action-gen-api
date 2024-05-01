package main

import (
	"flag"
	"fmt"
	"github.com/chamodshehanka/gh-action-gen-api/config"
	"github.com/chamodshehanka/gh-action-gen-api/server"
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
	//db.Init()
	server.Init()
}
