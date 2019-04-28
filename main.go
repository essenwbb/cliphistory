package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	globalDb = getDatabaseHandle()
	defer globalDb.Close()
	for {
		if ok := robotgo.AddEvents("c", "ctrl"); ok {
			recordClipToDatabase() //record to sqlite3
		}
	}
}
