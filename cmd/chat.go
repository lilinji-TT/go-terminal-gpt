package cmd

import (
	"GTG/config"
	"GTG/model"
	"GTG/pkg/gpt"
	utils "GTG/utils/functions"
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use: "chat",
}

var GlobMessages []model.Message

func Chat(cmd *cobra.Command, args []string) {


	if missConfig() {
		fmt.Println("Please set your config, url and api key. GTG config -u <your url> -k <your api key>")
		return
	}

	goos := runtime.GOOS // 获取操作系统

	fmt.Println("Welcome To Use GoTerminalGPT")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Print("User ~ % ")
		scanner.Scan()
		fmt.Println()
		userInput := scanner.Text()
		if userInput == "exit" {
			break
		}

		if userInput == "model" {
			config.SetModelName()
			continue
		}

		if userInput == "new chat" {
			fmt.Println(goos)
			switch goos {
			case "darwin":
				utils.OpenTerminal()
			case "linux":
				fmt.Println("Sorry! Linux is not supported.")
			case "windows":
				utils.OpenCmd()
			default:
				fmt.Printf("Unknown OS: %s\n", goos)
			}

			continue
		}

		fmt.Printf("%s ~ %% ", config.Model)
		gpt.GenerateStreamWithGPT(userInput, &GlobMessages, config.Model)
		fmt.Println()
	}
}

func missConfig() bool {
	url, key, err := config.ReadConfig()
	return err != nil || len(url) == 0 || len(key) == 0
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Run = Chat
}
