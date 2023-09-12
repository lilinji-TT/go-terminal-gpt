package cmd

import (
	"GTG/config"
	"GTG/model"
	"GTG/pkg/gpt"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use: "chat",
}

var GlobMessages []model.Message
func Chat(cmd *cobra.Command, args []string){

	isConfig := missConfig()

	if isConfig {
		fmt.Println("Please set your config, url and api key. GTG config -u <your url> -k <your api key>")
		return
	}

	fmt.Println("Welcome To Use GoTermiGPT")
	//创建读取用户输入的scanner
	scanner := bufio.NewScanner(os.Stdin)
	//循环读取
	for {
		fmt.Print("User ~ % ")
		scanner.Scan()
		userInput := scanner.Text()
		if (userInput == "exit") {
			break
		}
		fmt.Print("GPT ~ % ")
		gpt.GenerateStreamWithGPT(userInput, &GlobMessages)
		fmt.Println()
	}
}

func missConfig() bool {
	url, key, err := config.ReadConfig()
	return err != nil || len(url) == 0 || len(key) == 0
}

func init(){
	rootCmd.AddCommand(chatCmd)
	chatCmd.Run = Chat
}