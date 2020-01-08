package main

import "github.com/spf13/cobra"

func cmdInit() *cobra.Command {
	return &cobra.Command{
		Use: "init [command]",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
