package main

import (
	// go stdlib
	"os"
	"log"
	// juke subpackages
	"nu.hyperboreal.zone/juke/model/config"
)

func main() {
	// startup routine
	// check to make sure environment var JUKE_DATA_DIR is set;
	// die if it isn't
	if x := os.Getenv("JUKE_DATA_DIR"); x == "" {
		log.Fatal("Environment variable `JUKE_DATA_DIR` not set.")
		os.Exit(1)
	}
	// check if juke.toml exists
	if _, err := os.Stat(os.Getenv("JUKE_DATA_DIR") + "juke.toml");
			os.IsNotExist(err) {
		log.Fatal("Can't access config: ", err)
		os.Exit(1)
	}
	// load config from juke.toml

}
