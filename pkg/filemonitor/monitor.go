package filemonitor

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/clwg/eve-analyzer/pkg/model"
	"github.com/fsnotify/fsnotify"
)

func MonitorDirectory(dirPath string, pattern string, dataChan chan model.Event) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
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
				if event.Op&fsnotify.Create == fsnotify.Create {
					if matched, _ := filepath.Match(pattern, filepath.Base(event.Name)); matched {
						processFile(event.Name, dataChan)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func processFile(filePath string, dataChan chan model.Event) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var data model.Event
		if err := json.Unmarshal(scanner.Bytes(), &data); err != nil {
			log.Println("Error unmarshaling JSON line:", err)
			continue
		}
		dataChan <- data
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
}

func OpenFile(filePath string, dataChan chan model.Event) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var data model.Event
		if err := json.Unmarshal(scanner.Bytes(), &data); err != nil {
			log.Println("Error unmarshaling JSON line:", err)
			continue
		}
		dataChan <- data
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}

	close(dataChan)
}
