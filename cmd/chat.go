package cmd

import (
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


func init(){
	chatCmd.Run = Chat
}