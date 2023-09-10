package cmd

import (
	"GTG/pkg/prompt"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var promptCmd = &cobra.Command{
	Use: "prompt",
}


func init(){
	promptCmd.Run = func (cmd *cobra.Command, args []string){
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("You are setting prompt >> ")
		scanner.Scan()
		userInput := scanner.Text()
		prompt.SetPrompt(userInput)
	}
}