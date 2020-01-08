package main

import (
	"fmt"
	"github.com/glvd/cup/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func cmdRun() *cobra.Command {
	return &cobra.Command{
		Use: "run [command]",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.LoadJSON()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("Host:", config.Get().Host, "viperHost:", viper.GetString("host"))
		},
	}
}
