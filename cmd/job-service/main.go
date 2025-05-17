package main

import (
	"fmt"
	"log"

	"github.com/pm-cloudify/job-service/internal/config"
	"github.com/pm-cloudify/job-service/internal/service"
	"github.com/pm-cloudify/shared-libs/mb"
)

func main() {
	fmt.Println("launching...	")

	config.LoadConfigs()

	app_mb, err := mb.InitMessageBroker(
		config.AppConfigs.RMQ_Addr,
		config.AppConfigs.RMQ_User,
		config.AppConfigs.RMQ_Pass,
		config.AppConfigs.RMQ_Q_Name,
	)
	if err != nil {
		log.Fatalf("Failed to initialize message broker: %v", err)
	}
	defer app_mb.Close()

	err = mb.RunConsumer(app_mb, service.JobService)
	if err != nil {
		log.Fatalf("Failed to run consumer: %v", err)
	}
}
