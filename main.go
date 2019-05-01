package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/wuyueCreator/cliphistory/hotkey"
)

var (
	keys = map[int16]*hotkey.HotkeyAndEvent{
		1: {&hotkey.Hotkey{1, hotkey.ModShift + hotkey.ModCtrl, 'V'}, test},
	}
)

func main() {
	globalDb = getDatabaseHandle()
	defer globalDb.Close()

	hotkey.KeysEventMap = keys
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
