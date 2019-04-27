package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", "./dataSql.db")
	defer db.Close()
	checkErr(err)
	initDatabase(db)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initDatabase(db *sql.DB) {
	_, err := db.Query("select * from 'clipboard'")
	if err != nil {
		stmt, err := db.Prepare(`CREATE TABLE "clipboard"
			(
				"uid"        INTEGER PRIMARY KEY AUTOINCREMENT,
				"text"       VARCHAR(2000) NULL,
				"created"    DATE        NULL
			);
			`)
		checkErr(err)

		_, err = stmt.Exec()
		checkErr(err)

		_, err = db.Query("select * from 'clipboard'")
		checkErr(err)
	}
}
