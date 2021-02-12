package main

import (
	"flag"
	"fmt"
	"github.com/MBoegers/go-state-store/configuration"
	"github.com/MBoegers/go-state-store/controller"
	"github.com/MBoegers/go-state-store/datastore"
	"github.com/MBoegers/go-state-store/eventcache"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	// read configuration
	var configPath, err = readConfigPath()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var config *configuration.ConfigFile
	config, err = configuration.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// create events channels
	var newEventChannel = make(chan string)
	var publishChannel = make(chan []string)

	// init stores
	datastore.Init(newEventChannel)
	eventcache.Init(newEventChannel, publishChannel, config.Event.Time*time.Second)

	// spawn  servers
	var wg = new(sync.WaitGroup)
	wg.Add(2)
	var env = configuration.ReadEnv()

	go controller.SpawnEditCtl(config.Edit.Host, config.Edit.Port, wg)
	go controller.InitReadCtl(config.Read.Host, config.Read.Port,
		env.Cert, env.Key, publishChannel, wg)
	wg.Wait()
}

func readConfigPath() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	// ValidateConfigPath just makes sure, that the path provided is a file,
	// that can be read
	var s, err = os.Stat(configPath)
	if err != nil {
		return configPath, err
	}
	if s.IsDir() {
		return configPath, fmt.Errorf("'%s' is a directory, not a normal file", configPath)
	}
	// Return the configuration path
	return configPath, nil
}
