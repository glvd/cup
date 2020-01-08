package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const programName = `conversion`

var rootCmd = &cobra.Command{
	Use:     "conversion [command]",
	Short:   "call conversion command",
	Version: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("arguments [command] was not inputted")
	},
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 1,
}

func main() {
	rootCmd.AddCommand()
	e := rootCmd.Execute()
	if e != nil {
		panic(e)
	}
}

func viperRoot() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.BindPFlags()
	viper.BindFlagValues()
	err := viper.WriteConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
