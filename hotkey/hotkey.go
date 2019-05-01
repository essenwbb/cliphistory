package hotkey

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"
)

const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

type Hotkey struct {
	Modifiers int // Mask of modifiers
	KeyCode   int // Key code, e.g. 'A'
}

// String returns a human-friendly display name of the hotkey
// such as "Hotkey[Id: 1, Alt+Ctrl+O]"
func (h *Hotkey) String() string {
	mod := &bytes.Buffer{}
	if h.Modifiers&ModAlt != 0 {
		mod.WriteString("Alt+")
	}
	if h.Modifiers&ModCtrl != 0 {
		mod.WriteString("Ctrl+")
	}
	if h.Modifiers&ModShift != 0 {
		mod.WriteString("Shift+")
	}
	if h.Modifiers&ModWin != 0 {
		mod.WriteString("Win+")
	}
	return fmt.Sprintf("Hotkey[Id: %d, %s%c]", h.Id, mod, h.KeyCode)
}

type HotkeyAndEvent struct {
	Hotkey *Hotkey
	Event  func()
}

var (
	KeysEvent = make(map[int16]*HotkeyAndEvent)
)

func Run() {
	addHotkeys()
	listenEvent()
}

func addHotkeys() {
	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()

	reghotkey := user32.MustFindProc("RegisterHotKey")

	// Register hotkeys:
	for id, v := range KeysEvent {
		r1, _, err := reghotkey.Call(
			0, uintptr(id), uintptr(v.Hotkey.Modifiers), uintptr(v.Hotkey.KeyCode))
		if r1 == 1 {
			fmt.Println("Registered", v)
		} else {
			fmt.Println("Failed to register", v, ", error:", err)
		}
	}
}

type msg struct {
	HWND   uintptr
	UINT   uintptr
	WPARAM int16
	LPARAM int64
	DWORD  int32
	POINT  struct{ X, Y int64 }
}

func listenEvent() {
	user32 := syscall.MustLoadDLL("user32")
	defer user32.Release()
	peekmsg := user32.MustFindProc("PeekMessageW")

	for {
		var msg = &msg{}
		_, _, _ = peekmsg.Call(uintptr(unsafe.Pointer(msg)), 0, 0, 0, 1)

		// Registered id is in the WPARAM field:
		if id := msg.WPARAM; id != 0 {
			KeysEvent[id].Event()
		}
	}
}
