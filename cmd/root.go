package cmd

import (
	"GTG/model"
	"GTG/pkg/gpt"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:"GTG",
	Short:"A short description.",
	Long:"A longer description.",
}


var chatCmd = &cobra.Command{
	Use: "chat",
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var GlobMessages []model.Message
func chat(cmd *cobra.Command, args []string){
	//创建读取用户输入的scanner
	scanner := bufio.NewScanner(os.Stdin)
	//循环读取
	for {
		fmt.Print("User >>")
		scanner.Scan()
		userInput := scanner.Text()
		if (userInput == "exit") {
			break
		}
		fmt.Print("GPT >>")
		gpt.GenerateStreamWithGPT(userInput, &GlobMessages)
		fmt.Println()
	}
}

func init() {
	rootCmd.AddCommand(chatCmd)

	chatCmd.Run = chat
}