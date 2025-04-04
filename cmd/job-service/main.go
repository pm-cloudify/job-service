package main

import (
	"log"
	"os"

	"github.com/pm-cloudify/job-service/internal/service"
	"github.com/pm-cloudify/shared-libs/config_loader"
	"github.com/pm-cloudify/shared-libs/mb"
)

func main() {
	config_loader.LoadEnv("./configs")

	var (
		addr   = os.Getenv("RMQ_ADDR")
		user   = os.Getenv("RMQ_USER")
		pass   = os.Getenv("RMQ_PASS")
		q_name = os.Getenv("RMQ_Q_NAME")
	)

	app_mb, err := mb.InitMessageBroker(addr, user, pass, q_name)
	if err != nil {
		log.Fatalf("Failed to initialize message broker: %v", err)
	}
	defer app_mb.Close()

	err = mb.RunConsumer(app_mb, service.JobService)
	if err != nil {
		log.Fatalf("Failed to run consumer: %v", err)
	}
}
