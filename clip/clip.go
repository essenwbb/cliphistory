package clip

import (
	"github.com/atotto/clipboard"
	"io"
	"os"
	"strconv"
	"time"
)

func RecordClip(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}

	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	addIdentifier(w)
	_, _ = w.Write(append([]byte(text), '\n'))
}

func addIdentifier(w io.Writer) {
	currentTime := time.Now().Year()
	_, _ = w.Write(append([]byte(strconv.Itoa(int(currentTime))), '\n'))
}

func UpdateClipboard(text string) error {
	return clipboard.WriteAll(text)
}
