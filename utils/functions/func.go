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
	dir := GetCurrentPath()
	path := filepath.Join(dir, os.Args[0])
	cmd := exec.Command("cmd.exe", "/C", "start", "cmd.exe", "/K", path)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
