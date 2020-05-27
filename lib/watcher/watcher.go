package watcher

import (
	"log"
	"os"
	"time"
)

func WatchFile(filePath string, onChange func()) {
	go func(filePath string, onChange func()) {
		stat, err := os.Stat(filePath)
		if err != nil {
			log.Fatal(err)
		}

		size := stat.Size()
		modTime := stat.ModTime()

		for {
			newStat, err := os.Stat(filePath)
			if err != nil {
				log.Fatal(err)
			}

			if newStat.Size() != size || newStat.ModTime() != modTime {
				size = newStat.Size()
				modTime = newStat.ModTime()

				onChange()
			}

			time.Sleep(1 * time.Second)
		}
	}(filePath, onChange)
}
