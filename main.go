package main

import (
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/integrii/flaggy"
)

var (
	style = "none"
)

func main() {
	flaggy.String(&style, "s", "style", "the formatting style ('iso' / 'long-date' / 'time')")
	flaggy.Parse()

	now := time.Now()
	datestring := getDateString(now, style)

	toClipboard([]byte(datestring), runtime.GOOS)
}

func getDateString(now time.Time, style string) string {
	var datestring string
	switch style {
	case "iso":
		datestring = now.Format("2006-01-02")
	case "long-date":
		datestring = now.Format("January 2, 2006")
	case "time":
		datestring = now.Format("3:04 PM")
	default:
		datestring = now.String()
	}
	return datestring
}

func toClipboard(output []byte, arch string) {
	var copyCmd *exec.Cmd

	// Mac "OS"
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}
	// Linux
	if arch == "linux" {
		copyCmd = exec.Command("xclip", "-selection", "c")
	}

	in, err := copyCmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := in.Write([]byte(output)); err != nil {
		log.Fatal(err)
	}

	if err := in.Close(); err != nil {
		log.Fatal(err)
	}

	copyCmd.Wait()
}
