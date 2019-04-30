package main

import (
	"fmt"
	"github.com/wuyueCreator/cliphistory/hotkey"
)

var (
	keys = map[int16]*hotkey.HotkeyAndEvent{
		1: {&hotkey.Hotkey{1, hotkey.ModAlt + hotkey.ModCtrl, 'O'}, test},
	}
)

func main() {
	hotkey.KeysEventMap = keys
	hotkey.Run()
}

func test() {
	fmt.Println("hello world")
}
