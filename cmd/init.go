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
			cfg := &config.Config{
				Host: "127.0.0.1:6379",
			}
			config.Set(cfg)
			err := config.SaveJSON()
			if err != nil {
				log.Fatal(err)
				return
			}

		},
	}
}
