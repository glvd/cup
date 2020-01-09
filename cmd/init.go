package main

import (
	config2 "github.com/RichardKnop/machinery/v1/config"
	"github.com/glvd/cup/config"
	"github.com/spf13/cobra"
	"log"
)

func cmdInit() *cobra.Command {
	return &cobra.Command{
		Use: "init [command]",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := &config.Config{
				Broker:        "127.0.0.1:6379",
				QueueName:     "",
				ResultBackend: "",
				AMQP: config2.AMQPConfig{
					Exchange:         "1111",
					ExchangeType:     "",
					QueueDeclareArgs: nil,
					QueueBindingArgs: nil,
					BindingKey:       "",
					PrefetchCount:    0,
					AutoDelete:       false,
				},
			}
			config.Set(cfg)
			err := config.SaveConfig()
			if err != nil {
				log.Fatal(err)
				return
			}

		},
	}
}
