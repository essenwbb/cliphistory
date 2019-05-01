package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/wuyueCreator/cliphistory/hotkey"
)

var (
	keys = map[int16]*hotkey.HotkeyAndEvent{
		1: {&hotkey.Hotkey{hotkey.ModShift + hotkey.ModCtrl, 'V'}, test},
	}
)

func main() {
	globalDb = getDatabaseHandle()
	defer globalDb.Close()

	hotkey.KeysEvent = keys
	go hotkey.Run()

	for {
		if ok := robotgo.AddEvents("c", "ctrl"); ok {
			recordClipToDatabase() //record to sqlite3
		}
	}
}

func test() {
	fmt.Println("hello world")
}
