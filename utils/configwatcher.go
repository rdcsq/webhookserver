package utils

import (
	"log"
	"webhookserver/structs"

	"github.com/fsnotify/fsnotify"
)

func WatchConfig() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					if structs.ParseConfig() {
						log.Println("Config file modified, reloading")
					} else {
						log.Println("Error parsing config file. API will continue running with the previous config")
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				panic(err)
			}
		}
	}()

	err = watcher.Add(structs.Env.ConfigPath)
	if err != nil {
		panic(err)
	}

	return watcher
}
