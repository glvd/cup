package main

import (
	"fmt"
	"github.com/glvd/cup"
	"github.com/glvd/cup/config"
	"github.com/glvd/cup/service"
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

			s := service.NewService(*config.Get())
			s.NewWorker()
			err = s.Register("slice", cup.TaskSlice)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = s.HandleWorker()
			if err != nil {
				log.Fatal(err)
				return
			}
		},
	}
}
