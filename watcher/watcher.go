package watcher

import (
	"ivnvMkhl/live-server/logger"
	"log"

	"github.com/fsnotify/fsnotify"
)

type Event = fsnotify.Event

func Subscribe(path string, cb func(fsnotify.Event)) (dispose func()) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					cb(event)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Log(true, err.Error())
			}
		}
	}()

	// Add a path.
	err = watcher.Add(path)
	if err != nil {
		logger.Fatal("Not found wath dir", err)
	}
	return func() { watcher.Close() }
}
