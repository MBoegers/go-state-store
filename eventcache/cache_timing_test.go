package eventcache

import (
	"testing"
	"time"
)

func TestSingle(t *testing.T) {
	//region<Setup>
	var newChan = make(chan string)
	var publishChan = make(chan []string)
	Init(newChan, publishChan, 1*time.Second)
	var actualEvents int

	//spawn receiver
	go func() {
		select {
		case events := <-publishChan:
			actualEvents = len(events)
		}
	}()
	//send new
	var timer = time.NewTimer(1500 * time.Millisecond)
	newChan <- "TestEvent"
	//wait
	<-timer.C
	//assert
	if actualEvents != 1 {
		t.Errorf("Should have received one update, got: %v", actualEvents)
	}
}

func TestThree(t *testing.T) {
	//region<Setup>
	var newChan = make(chan string)
	var pubChan = make(chan []string)
	Init(newChan, pubChan, 1*time.Second)
	var actualEvents int

	//spawn receiver
	go func() {
		select {
		case events := <-pubChan:
			actualEvents = len(events)
		}
	}()
	//send new
	var timer = time.NewTimer(1500 * time.Millisecond)
	newChan <- "firstEvent"
	newChan <- "secondEvent"
	newChan <- "thirdEvent"
	//wait
	<-timer.C
	//assert
	if actualEvents != 3 {
		t.Errorf("Should have received three update, got: %v", actualEvents)
	}
}
