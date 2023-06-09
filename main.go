package main

import (
	// go stdlib
	"os"
	"log"

	"hyperboreal.zone/juke/model/config"
)

func main() {
	// startup routine
	// check if juke.toml exists
	if _, err := os.Stat("juke.toml"); os.IsNotExist(err) {
		log.Fatal("Can't access config:", err)
	}
	// load config from juke.toml

}
