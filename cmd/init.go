package main

import (
	"github.com/glvd/cup/config"
	"github.com/spf13/cobra"
	"log"
)

func cmdInit() *cobra.Command {
	return &cobra.Command{
		Use: "init [command]",
		Run: func(cmd *cobra.Command, args []string) {
			config.Set(config.Default())
			err := config.SaveConfig()
			if err != nil {
				log.Fatal(err)
				return
			}

		},
	}
}
