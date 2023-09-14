package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

const (
	BaseURL   = ""
	OpenaiKey = ""
)

var Model string = "gpt-3.5-turbo"

var modelOptions = []*survey.Question{
	{
		Name: "model",
		Prompt: &survey.Select{
			Message: "Which model do you want to use?",
			Options: []string{
				"gpt-3.5-turbo",
				"gpt-3.5-turbo-16k",
				"gpt-4",
			},
		},
		Validate: survey.Required,
	},
}

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

func ReadConfig() (string, string, error) {
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

func SetModelName() {
	err := survey.Ask(modelOptions, &Model)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Now Model is ", Model)
}
