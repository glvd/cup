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

	cmd := &cobra.Command{
		Use: "run [command]",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.Load()
			if err != nil {
				log.Fatal(err)
				return
			}
			cfg := config.Get()
			fmt.Printf("config:%+v\n", cfg)
			cup.TaskInit(config.Get())
			s := service.NewService(*cfg)
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
	//cmd.Flags().StringP("ffmpeg", "f", "./bin", "set the ffmpeg command path")
	return cmd
}
