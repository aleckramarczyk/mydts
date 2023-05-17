package utils

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"math"
	"strings"
	"sync"
	"time"
)

func WatchForEvents() {
	// Create new watcher to watch for write events in notable directories. We do this because
	// Intergraph will create archive files whenever a sign on or sign off is made, which cues
	// our program to send a request to update the server
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("ERROR creating new watcher: %s", err)
	}
	defer w.Close()

	//Start listening for events
	go deduplicatingLoop(w)

	paths := [...]string{AgentConfig.FromClientGPSPath, AgentConfig.FromClientAuthorizePath, AgentConfig.FromClientSignOffPath}

	for _, p := range paths {
		err = w.Add(p)
		if err != nil {
			log.Fatalf("Error watching path: %s", err)
		}
	}

	log.Print("Watching for file events")
	<-make(chan struct{}) // Block indefinitely
}

func deduplicatingLoop(w *fsnotify.Watcher) {
	// A single write can generate many Write events, so to prevent sending several requests
	// when only one real event occurs we need to deduplicate these write events. Here this is
	// done by waiting a short time for more write events and resetting the wait period for
	// every new event. See the fsnotify dedup examples for more info.
	var (
		// Wait 1 second for new events
		waitFor = time.Second

		// Keep track of the timers on file objects
		mu     sync.Mutex
		timers = make(map[string]*time.Timer)

		// Function that is called when an event has been deduplicated
		determineEventType = func(e fsnotify.Event) {
			mu.Lock()
			delete(timers, e.Name)
			// event.Name is the path of the file that was written to
			mu.Unlock()

			switch {
			/*
				case strings.Contains(e.String(), "GPS"):
					entities.UnitInformation.UpdateGPSInformation(e.String())
			*/
			case strings.Contains(e.String(), "SignOff"):
				UnitInformation.UpdateUnitPropertiesOnSignOff(e.Name)
			case strings.Contains(e.String(), "Authorize"):
				UnitInformation.UpdateUnitPropertiesOnSignOn(e.Name)
			}
		}
	)

	for {
		select {
		case err, ok := <-w.Errors:
			if !ok { // Channel was closed
				return
			}
			log.Println(err)
		case e, ok := <-w.Events:
			// We are only looking for create and write events, so ignore everything else
			if !e.Has(fsnotify.Create) && !e.Has(fsnotify.Write) {
				continue
			}

			// Get timer
			mu.Lock()
			t, ok := timers[e.Name]
			mu.Unlock()

			// No timer yet, so create one
			if !ok {
				t = time.AfterFunc(math.MaxInt64, func() { determineEventType(e) })
				t.Stop()

				mu.Lock()
				timers[e.Name] = t
				mu.Unlock()
			}

			// Reset the timer for this path
			t.Reset(waitFor)
		}
	}
}
