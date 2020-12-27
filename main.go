package main

import (
	"go-state-publisher/controller"
	"go-state-publisher/datastore"
	"go-state-publisher/eventcache"
	"sync"
	"time"
)

func main() {
	var newEventChannel = make(chan string)
	var publishChannel = make(chan []string)

	datastore.Init(newEventChannel)
	eventcache.Init(newEventChannel, publishChannel, 5*time.Second)

	var wg = new(sync.WaitGroup)
	wg.Add(2)

	go controller.SpawnEditCtl("127.0.0.1", 8081, wg)
	go controller.InitReadCtl(8080, publishChannel, wg)

	wg.Wait()
}
