package notify

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func Notify(filename string, onUpdate func()) (err error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Printf("%s updated\n", filename)
					onUpdate()
				}
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()
	err = watcher.Add(filename)
	if err != nil {
		return
	}
	<-done
	return
}
