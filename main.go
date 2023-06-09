package main

import (
	// go stdlib
	"os"
	"log"
	// juke subpackages
	"hyperboreal.zone/juke/model/config"
)

func main() {
	// startup routine
	// check if juke.toml exists
	if _, err := os.Stat("juke.toml"); os.IsNotExist(err) {
		log.Fatal("Can't access config:", err)
		os.Exit(1)
	}
	// load config from juke.toml

}
