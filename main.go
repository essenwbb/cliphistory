package main

import (
	"cliphistory/clip"
	"database/sql"
	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var globalDb *sql.DB

func main() {
	globalDb = getDatabaseHandle()
	defer globalDb.Close()

	for {
		if ok := robotgo.AddEvents("c", "ctrl"); ok {
			recordClipToFile()
			//record to sqlite
			recordClipToDatabase()
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

func recordClipToDatabase() {
	text, err := clipboard.ReadAll()
	checkErr(err)
	//插入数据
	stmt, err := globalDb.Prepare("INSERT INTO clipboard (text, created) VALUES (?, ?)")
	checkErr(err)

	_, err = stmt.Exec(text, time.Now().Format("2006-01-02"))
	checkErr(err)
}
