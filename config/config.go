package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	BaseURL   = ""
	OpenaiKey = ""
)


func WriteConfig(url string, key string) error {
	//获取当前用户主目录
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	data := url + " " + key
	configPath := filepath.Join(home, ".config_gtg")
	err = os.WriteFile(configPath, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}
	return nil
}

func ReadConfig()(string, string, error) {
	//获取当前用户主目录
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(home, ".config_gtg")
	// 读取文件内容
    data, err := os.ReadFile(configPath)
	if err != nil {
        return "", "", err
    }
	lines := strings.Split(string(data), " ")

	return lines[0], lines[1], nil
}