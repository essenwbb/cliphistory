package main

import (
	"fmt"
	"github.com/wuyueCreator/cliphistory/hotkey"
)

var (
	keys = map[int16]*hotkey.HotkeyAndEvent{
		1: {&hotkey.Hotkey{hotkey.ModAlt + hotkey.ModCtrl, 'O'}, test},
		2: {&hotkey.Hotkey{hotkey.ModCtrl, 'V'}, test2},
	}
)

func main() {
	hotkey.KeysEvent = keys
	hotkey.Run()
}

func test() {
	fmt.Println("hello world")
}

func test2() {
	fmt.Println("hello world 2")
}
