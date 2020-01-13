package main

import (
	"encoding/json"
	"fmt"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/glvd/cup"
	"github.com/glvd/cup/config"
	"github.com/glvd/cup/service"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func cmdSend() *cobra.Command {
	cfg, err := json.Marshal(config.DefaultSliceConfig())
	if err != nil {
		return nil
	}
	signature := &tasks.Signature{
		Name: "slice",
		Args: []tasks.Arg{
			{
				Type:  "[]byte",
				Value: cfg,
			},
		},
	}

	return &cobra.Command{
		Use: "send [command]",
		Run: func(cmd *cobra.Command, args []string) {
			err := config.Load()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("config:%+v\n", config.Get())

			s := service.NewService(*config.Get())
			rlt, err := s.Send(signature)
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Printf("result:%+v\n", rlt)
			for !rlt.GetState().IsCompleted() {
				log.Println("running", rlt.GetState().State)
				time.Sleep(1 * time.Second)
			}
			for _, result := range rlt.GetState().Results {
				log.Println("result", result)
				var f cup.Fragment
				err := json.Unmarshal([]byte(result.Value.(string)), &f)
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	}
}
