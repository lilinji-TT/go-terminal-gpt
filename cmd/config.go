package cmd

import (
	"GTG/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var URL string

var ApiKey string

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Set base config, such as base url and openai key",
	Run: func(cmd *cobra.Command, args []string) {
		URL, _ := cmd.Flags().GetString("url")
		ApiKey, _ := cmd.Flags().GetString("key")
		if err := config.WriteConfig(URL, ApiKey); err != nil {
			fmt.Println("Can not write config file:", err)
			os.Exit(1)
		} else {
			fmt.Println("Config file has been written.")
		}
	},
}

func init(){
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVarP(&URL, "url", "u", "", "base url")
	configCmd.Flags().StringVarP(&ApiKey, "key", "k", "", "api key for openai authentication")
}