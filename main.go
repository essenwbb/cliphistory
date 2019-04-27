package main

import (
	"cliphistory/clip"
	"github.com/go-vgo/robotgo"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	for {
		if ok := robotgo.AddEvents("c", "ctrl"); ok {
			recordClipToFile()
			//record to sqlite

		}
	}
}

func recordClipToFile() {
	f, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	clip.RecordClip(f)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
