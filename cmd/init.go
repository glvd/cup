package main

import (
	"github.com/glvd/cup/config"
	"github.com/spf13/cobra"
	"log"
)

func cmdInit() *cobra.Command {
	path := ""

	cmd := &cobra.Command{
		Use: "init [command]",
		Run: func(cmd *cobra.Command, args []string) {
			config.Set(config.Default())
			config.DefaultConfigPath = path
			err := config.Save()
			if err != nil {
				log.Fatal(err)
				return
			}

		},
	}
	cmd.Flags().StringP("path", "p", ".", "set the output config path")
	return cmd
}
