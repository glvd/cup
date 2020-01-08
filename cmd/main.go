package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

const programName = `cup`

var rootCmd = &cobra.Command{
	Use:     "cup [command]",
	Short:   "call cup command",
	Version: "v0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("arguments [command] was not inputted")
	},
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 1,
}

func main() {
	rootCmd.AddCommand(cmdInit(), cmdRun())
	e := rootCmd.Execute()
	if e != nil {
		panic(e)
	}
}
