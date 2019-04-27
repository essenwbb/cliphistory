package main

import (
	"cliphistory/clip"
	"github.com/go-vgo/robotgo"
	"log"
	"os"
)

func main() {
	for {
		if ok := robotgo.AddEvents("c", "ctrl"); ok {
			recordClipToFile()
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
