package main

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/glvd/cup"
	"github.com/glvd/cup/config"
	"github.com/glvd/cup/service"
	"github.com/spf13/cobra"
)

func cmdSend() *cobra.Command {
	signature := &tasks.Signature{
		Name: "send",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

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
			err = s.Register("slice", cup.Slice)
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
