package main

import (
	_ "embed"
	"log"
	"os"
	"runtime"

	"gioui.org/app"
)

//go:embed fonts/TerminessNerdFont-Bold.ttf
var fontData []byte

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		os.Setenv("HOME", path)
	}
	if runtime.GOOS == "windows" {
		os.Setenv("APPDATA", path)
	}
	go func() {
		application := NewApp(600, 900, "Утилита НИС")
		err := application.Run()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
