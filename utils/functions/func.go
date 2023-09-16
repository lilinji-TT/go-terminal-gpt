package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func OpenBrowse() {

}

func GetCurrentPath() string {
	curdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return curdir
}

func OpenCmd() {
	fmt.Println("Open Cmd")
	dir := GetCurrentPath()
	path := filepath.Join(dir, os.Args[0])
	cmd := exec.Command("cmd.exe", "/C", "start", "cmd.exe", "/K", path)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func OpenTerminal() {
	fmt.Println("Open Terminal")
	script := `tell application "Terminal" to do script "GTG chat"`
	cmd := exec.Command("osascript", "-e", script)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to start terminal:", err)
	}
}
