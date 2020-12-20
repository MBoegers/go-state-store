package eventcache

import (
	"sync"
	"time"
)

var newChan chan string
var notifyChan chan []string
var sendSchedule *time.Ticker
var eventQueue []string
var eventLock sync.Mutex

func Init(newEventChan chan string, publishChan chan []string, sendTime time.Duration) {
	newChan = newEventChan
	notifyChan = publishChan
	sendSchedule = time.NewTicker(sendTime)
	eventQueue = make([]string, 0)

	go handleMessages()
}

//region<Handling>
func handleMessages() {
	for {
		select {
		case <-sendSchedule.C:
			sendEvents()
		case event := <-newChan:
			addEvent(event)
		}
	}
}

// append a new events to the event queue concurrently
func addEvent(event string) {
	eventLock.Lock()
	defer eventLock.Unlock()

	eventQueue = append(eventQueue, event)
}

// send all events out and clear event queue
func sendEvents() {
	eventLock.Lock()
	defer eventLock.Unlock()

	var payload = make([]string, len(eventQueue))
	copy(payload, eventQueue)
	eventQueue = make([]string, 0)
	notifyChan <- payload
}

//endregion
