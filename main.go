package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	now := time.Now()
	year, month, day := now.Date()
	datestring := fmt.Sprintf("%04d-%02d-%02d", year, month, day)

	toClipboard([]byte(datestring), runtime.GOOS)
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
