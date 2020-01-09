package main

import (
	"fmt"
	"github.com/glvd/cup/config"
	"github.com/spf13/cobra"
	"log"
)

func cmdRun() *cobra.Command {
	return &cobra.Command{
		Use: "run [command]",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.LoadConfig()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("config:%+v\n", config.Get())
		},
	}
}
