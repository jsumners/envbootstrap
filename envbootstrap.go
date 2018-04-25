package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/jsumners/envbootstrap/lib"
)

func main() {
	var config *viper.Viper
	var err error

	configFilePath := flag.String("config", "", "path to configuration file (Required)")
	flag.Parse()

	config, err = lib.Config(*configFilePath)
	if err != nil {
		fmt.Printf("Unable to read configuration %s: %s", *configFilePath, err)
		os.Exit(1)
	}

	fmt.Print(config)
}
