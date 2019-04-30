Example code:

```go
package main

import (
	"fmt"
	"github.com/wuyueCreator/cliphistory/hotkey"
)

var (
	keys = map[int16]*hotkey.HotkeyAndEvent{
		1: {&hotkey.Hotkey{1, hotkey.ModAlt + hotkey.ModCtrl, 'O'}, test},
		2: {&hotkey.Hotkey{2, hotkey.ModCtrl, 'V'}, test2},
	}
)

func main() {
	hotkey.KeysEventMap = keys
	hotkey.Run()
}

func test() {
	fmt.Println("hello world")
}

func test2() {
	fmt.Println("hello world 2")
}
```