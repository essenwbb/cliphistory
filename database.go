package main

import (
	"database/sql"
	"fmt"
	"github.com/atotto/clipboard"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var (
	globalDb *sql.DB
	lastUid  int
)

func init() {
	db := getDatabaseHandle()
	defer db.Close()

	checkDatabase(db)
	echoNew(db)
}

func getDatabaseHandle() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./dataSql.db")
	checkErr(err)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkDatabase(db *sql.DB) {
	_, err := db.Query("select * from 'clipboard'")
	if err != nil {
		stmt, err := db.Prepare(`CREATE TABLE "clipboard"
			(
				"uid"        INTEGER PRIMARY KEY AUTOINCREMENT,
				"text"       VARCHAR(2000) NULL,
				"created"    TIMESTAMP     NULL
			);
			`)
		checkErr(err)

		_, err = stmt.Exec()
		checkErr(err)

		_, err = db.Query("select * from 'clipboard'")
		checkErr(err)
	}
}

func listAll(db *sql.DB) {
	rows, err := db.Query("SELECT uid, text, created FROM clipboard")
	checkErr(err)

	for rows.Next() {
		var uid int
		var text string
		var created time.Time
		err = rows.Scan(&uid, &text, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(text)
		fmt.Println(time.Unix(created.Unix(), 0).Format("2006-1-2 15:04:05"))
	}
}

func echoNew(db *sql.DB) {
	rows, err := db.Query("select * from clipboard order by uid DESC limit 1;")
	checkErr(err)

	for rows.Next() {
		var uid int
		var text string
		var created time.Time
		err = rows.Scan(&uid, &text, &created)
		checkErr(err)

		lastUid = uid
		fmt.Println(uid)
		fmt.Println(text)
		fmt.Println(time.Unix(created.Unix(), 0).Format("2006-1-2 15:04:05"))
	}
}

func toppingClipById(db *sql.DB, uid int) {
	stmt, err := db.Prepare("update clipboard set uid=? where uid==?")
	checkErr(err)

	res, err := stmt.Exec(lastUid+1, uid)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)
}

func recordClipToDatabase() {
	text, err := clipboard.ReadAll()
	checkErr(err)

	stmt, err := globalDb.Prepare("INSERT INTO clipboard (text, created) VALUES (?, ?)")
	checkErr(err)

	_, err = stmt.Exec(text, time.Now().UnixNano()/1e6)
	checkErr(err)

	echoNew(globalDb)
}
